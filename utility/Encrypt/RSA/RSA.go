package RSA

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	Config "github.com/Thenecromance/OurStories/utility/config"
	"os"
)

const (
	privateKeyFile = "setting/private.pem"
	publicKeyFile  = "setting/public.pem"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	files      filePath
)

// TODO: change RSA to implement the Encrypt interface

type filePath struct {
	PrivateKeyFile string `ini:"privateKeyFile"`
	PublicKeyFile  string `ini:"publicKeyFile"`
}

func init() {
	files.PrivateKeyFile = privateKeyFile
	files.PublicKeyFile = publicKeyFile
	err := Config.LoadToObject("Encrypt", &files)
	if err != nil {
		return
	}

	if !allPemExists() {
		GenerateKey()
	} else {
		load()
	}

}

func allPemExists() bool {
	// check public key and private key files
	return fileExists(files.PrivateKeyFile) && fileExists(files.PublicKeyFile)
}
func fileExists(filename string) bool {
	_, err :=
		os.Stat(filename)
	return err == nil
}

// GenerateKey will generate the public and private key
func GenerateKey() error {
	//generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	publicKey = &privateKey.PublicKey

	//dump private key to file
	err = dumpPrivateKeyFile(privateKey, files.PrivateKeyFile)
	if err != nil {
		return err
	}

	//dump public key to file
	err = dumpPublicKeyFile(&privateKey.PublicKey, files.PublicKeyFile)
	if err != nil {
		return err
	}

	return nil
}

func dumpPrivateKeyFile(privatekey *rsa.PrivateKey, filename string) error {
	var keybytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keybytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func dumpPublicKeyFile(publickey *rsa.PublicKey, filename string) error {
	keybytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keybytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func dumpPrivateKeyBase64(privatekey *rsa.PrivateKey) (string, error) {
	var keybytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)

	keybase64 := base64.StdEncoding.EncodeToString(keybytes)
	return keybase64, nil
}

func dumpPublicKeyBase64(publickey *rsa.PublicKey) (string, error) {
	keybytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", err
	}

	keybase64 := base64.StdEncoding.EncodeToString(keybytes)
	return keybase64, nil
}

// LoadPrivateKeyFile will load the private key from the file
func loadPrivateKeyFile(filename string) (*rsa.PrivateKey, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
func loadPublicKeyFile(filename string) (*rsa.PublicKey, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PublicKey), nil
}

func load() {
	var err error
	privateKey, err = loadPrivateKeyFile(files.PrivateKeyFile)
	if err != nil {
		GenerateKey()
		privateKey, err = loadPrivateKeyFile(files.PrivateKeyFile)
	}
	publicKey, err = loadPublicKeyFile(files.PublicKeyFile)
	if err != nil {
		GenerateKey()
		publicKey, err = loadPublicKeyFile(files.PublicKeyFile)
	}
}

func Encrypt(plaintext string) (string, error) {
	//encrypt
	cipherbytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plaintext))
	if err != nil {
		return "", err
	}
	//encode base64
	return base64.StdEncoding.EncodeToString(cipherbytes), nil
}

func Decrypt(ciphertext string) (string, error) {
	//decode base64
	cipherbytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	//decrypt
	plainbytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherbytes)
	if err != nil {
		return "", err
	}
	return string(plainbytes), nil
}
