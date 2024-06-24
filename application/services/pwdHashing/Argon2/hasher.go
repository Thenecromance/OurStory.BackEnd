package Argon2

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"github.com/Thenecromance/OurStories/application/services/pwdHashing"
	"github.com/Thenecromance/OurStories/utility/log"
	"golang.org/x/crypto/argon2"
)

type argon2or struct {
	setting *Setting
}

// randomSalt generates a random salt, which is controlled by the RandomSaltLen in the Setting
func (a *argon2or) randomSalt() (salt []byte) {
	salt = make([]byte, a.setting.RandomSaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return
}

func (a *argon2or) Hash(password string) (hashed string, salted string) {
	salt_ := a.randomSalt()
	argonHashed := argon2.IDKey(
		[]byte(password),
		salt_,
		a.setting.Time,
		a.setting.Memory,
		uint8(a.setting.Threads),
		a.setting.KeyLen)
	hashed = base64.StdEncoding.EncodeToString(argonHashed)
	salted = base64.StdEncoding.EncodeToString(salt_)
	return
}

func (a *argon2or) Verify(inputPwd, hashedPwd, salt string) bool {
	salt_, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		log.Error(err)
		return false
	}
	hashedPwd_, err := base64.StdEncoding.DecodeString(hashedPwd)
	if err != nil {
		log.Error(err)
		return false
	}
	argonHashed := argon2.IDKey(
		[]byte(inputPwd),
		salt_,
		a.setting.Time,
		a.setting.Memory,
		uint8(a.setting.Threads),
		a.setting.KeyLen)
	return bytes.Equal(argonHashed, hashedPwd_)
}

func New() pwdHashing.PwdHasher {
	return &argon2or{
		setting: newSetting(),
	}
}
