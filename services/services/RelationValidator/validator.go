package RelationValidator

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"github.com/Thenecromance/OurStories/utility/cache/lru"
	"github.com/Thenecromance/OurStories/utility/log"
	"time"
)

const (
	ExpireTime = time.Hour * 1
)

type relationCache struct {
	UserId       int
	RelationType int
	Idx          int
	Stamp        int64 //WARNING: this field use time.Now().UnixNano() not time.Now().Unix()
}

/*
should I put the data into the database? or just keep it in memory?
*/
type validator struct {
	cache *lru.Cache
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
func (v *validator) GenerateToken(userID int, relationType int, idx int) (string, error) {

	cache := &relationCache{
		UserId:       userID,
		RelationType: relationType,
		Idx:          idx,
		Stamp:        time.Now().UnixNano(),
	}
	hash := generateHash(cache)
	v.cache.Add(hash, cache, time.Now().Add(ExpireTime)) // just added it to the cache and give it an expiration time
	return hash, nil
}

func (v *validator) GetTokenInfo(token string) (userID int, relationType int, err error) {
	cache, ok := v.cache.Get(token)
	if !ok {
		return 0, 0, fmt.Errorf("token not found")
	}
	rc, ok := cache.(*relationCache)
	if !ok {
		return 0, 0, fmt.Errorf("token not found")
	}
	return rc.UserId, rc.RelationType, nil
}

func New() IRelationValidator {
	return &validator{
		cache: lru.New(0),
	}
}
