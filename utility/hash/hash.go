package hash

import (
	"crypto/sha256"
	"fmt"
)

var (
	Salt = "M0nk3y_Cl!ck3r!"
)

/*
TODO: this package should be deprecated after the scrypt package is finished
*/

// salted will combine the origin string and the salt string
//
// Deprecated: this function has been deprecated, use the Encrypt package instead
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

// Hash will hash the origin string with the salt
//
// Deprecated: this function has been deprecated, use the Encrypt package instead
func Hash(origin string) string {
	hash := sha256.Sum256([]byte(salted(origin, Salt)))
	return fmt.Sprintf("%x", hash)
}

// HashWithSalt will hash the origin string with the salt
//
// Deprecated: this function has been deprecated, use the Encrypt package instead
func HashWithSalt(origin, salt string) string {
	hash := sha256.Sum256([]byte(salted(origin, salt)))
	return fmt.Sprintf("%x", hash)
}

// Compare will hash the origin string with the salt
//
// Deprecated: this function has been deprecated, use the Encrypt package instead
func Compare(origin, hashed string) bool {
	return Hash(origin) == hashed
}
