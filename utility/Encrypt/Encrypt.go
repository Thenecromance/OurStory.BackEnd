package Encrypt

// Encryptor is an interface that defines the methods that an encryptor should implement
// this is useful for user authentication, where the password is hashed and stored in the database
type Encryptor interface {
	// Hash takes a password and returns a hashed password and a salt used to hash the password
	// due to the nature of the scrypt algorithm, the salt is generated randomly, so it is not necessary to provide it
	Hash(password string) (hashed string, salt string)
	// Verify takes a password, a hashed password and a salt and returns true if the password matches the hashed password
	Verify(inputPwd, hashedPwd, salt string) bool
}
