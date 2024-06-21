package hash

import (
	"crypto/sha256"
	"fmt"
)

var (
	Salt = "M0nk3y_Cl!ck3r!"
)

func salted(origin string, salt string) (saltedStr string) {
	oLen := len(origin)
	SLen := len(salt)

	var maxLen int
	if oLen > SLen {
		maxLen = oLen
	} else {
		maxLen = SLen
	}
	for i := 0; i < maxLen; i++ {
		if i < oLen {
			saltedStr += string(origin[i])
		}
		if i < SLen {
			saltedStr += string(salt[i])
		}
	}
	return
}

func Hash(origin string) string {
	hash := sha256.Sum256([]byte(salted(origin, Salt)))
	return fmt.Sprintf("%x", hash)
}

func HashWithSalt(origin, salt string) string {
	hash := sha256.Sum256([]byte(salted(origin, salt)))
	return fmt.Sprintf("%x", hash)
}

func Compare(origin, hashed string) bool {
	return Hash(origin) == hashed
}
