package RelationValidator

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/cache/redisCache"
	"github.com/Thenecromance/OurStories/utility/log"
)

const (
	ExpireTime = time.Hour * 1
)

const (
	prefix          = "relation"
	dbUserIdToToken = "user_id_to_token"
	dbTokenToUser   = "token_to_user"
)

type relationCache struct {
	UserId       int64 `json:"user_id,omitempty"`
	RelationType int   `json:"relation_type,omitempty"`
	Idx          int   `json:"idx,omitempty"`
	Stamp        int64 `json:"stamp,omitempty"` //WARNING: this field use time.Now().UnixNano() not time.Now().Unix()
}

/*
should I put the data into the database? or just keep it in memory?
*/
type validator struct {
	cache Interface.ICache // key:hashed token, value:relationCache
}

func generateHash(cache *relationCache) string {
	log.Debug("hashing ", cache)
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(cache) // ignore this error, if don't change any code, this will never error
	if err != nil {
		log.Error("error while encoding cache ", err)
	}
	return fmt.Sprintf("%x", sha1.Sum(buf.Bytes()))
}

func (v *validator) setToken(cache *relationCache, token string) {
	buf, err := json.Marshal(cache)
	if err != nil {
		return
	}
	// userId--->token
	// token--->UserCache

	v.cache.Prefix(dbTokenToUser)
	v.cache.Set(token, string(buf), ExpireTime)
	v.cache.Prefix(dbUserIdToToken)
	v.cache.Set(strconv.FormatInt(cache.UserId, 10), token, ExpireTime)
}

func (v *validator) getUser(token string) (rc *relationCache, err error) {
	v.cache.Prefix(dbTokenToUser)
	buf, err := v.cache.Get(token)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(buf.(string)), rc)
	if err != nil {
		return nil, err
	}
	return
}
func (v *validator) getToken(userId string) (rc *relationCache, err error) {
	v.cache.Prefix(dbUserIdToToken)
	token, err := v.cache.Get(userId)
	if err != nil {
		return nil, err
	}
	return v.getUser(token.(string))
}

func (v *validator) GenerateToken(userID int64, relationType int, idx int) (string, error) {

	cache := &relationCache{
		UserId:       userID,
		RelationType: relationType,
		Idx:          idx,
		Stamp:        time.Now().UnixNano(),
	}
	hash := generateHash(cache)
	v.setToken(cache, hash)
	return hash, nil
}

func (v *validator) GetTokenInfo(token string) (userID int64, relationType int, err error) {
	cache, err := v.getUser(token)
	return cache.UserId, cache.RelationType, err
}

func New() IRelationValidator {
	v := &validator{
		cache: redisCache.NewCache(),
	}
	v.cache.Prefix(prefix)
	return v
}
