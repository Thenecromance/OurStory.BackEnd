package Scrypt

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"github.com/Thenecromance/OurStories/services/services/pwdHashing"
	"github.com/Thenecromance/OurStories/utility/log"
	"golang.org/x/crypto/scrypt"
)

type scryptor struct {
	cfg *Setting
}

// randomSalt generates a random salt, which is controlled by the RandomSaltLen in the Setting
func (s *scryptor) randomSalt() (salt []byte) {
	salt = make([]byte, s.cfg.RandomSaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return
}

// generateKeyWithSalt generates a key with a given password and salt
func (s *scryptor) generateKeyWithSalt(password, salt []byte) (key []byte) {
	key, err := scrypt.Key(password, salt, s.cfg.CostFactor, s.cfg.BlockSizeFactor, s.cfg.ParallelizationFactor, s.cfg.KeyLen)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return
}

// Hash generates a hashed password and a salt for a given password
//
// @param: password: the password which need to be hashed
//
// @return: a base64 encoded string of the hashed password and the salt
func (s *scryptor) Hash(password string) (key, salt string) {
	saltBuffer := s.randomSalt()
	keyBuffer := s.generateKeyWithSalt([]byte(password), saltBuffer)

	key = base64.StdEncoding.EncodeToString(keyBuffer)
	salt = base64.StdEncoding.EncodeToString(saltBuffer)
	log.Debugf("key: %s salt: %s", key, salt)
	return
}

// Verify verifies a password with a given hash and salt
//
// @param: password: the password to verify
//
// @param: hash: the hashed password to verify against
//
// @param: salt: the salt to verify against, password need to be hashed with the same salt to be verified
//
// @return: true if the password is correct, false otherwise
func (s *scryptor) Verify(password, hash, salt string) bool {
	keyBuffer, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		log.Error(err)
		return false
	}
	saltBuffer, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		log.Error(err)
		return false
	}

	newKey := s.generateKeyWithSalt([]byte(password), saltBuffer)
	return bytes.Equal(newKey, keyBuffer)

}

func New() pwdHashing.PwdHasher {
	return &scryptor{
		cfg: newConfig(),
	}
}
