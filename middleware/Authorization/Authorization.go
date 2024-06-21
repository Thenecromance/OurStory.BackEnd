package Authorization

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	AuthObject = "UserData"
	EncryptKey = "M0nkey_Cl1cker"
)

// ITokenSigner is an interface for SignToken to users
type ITokenSigner interface {
	// Server need to provide a user claim to sign a token for client
	SignTokenToUser(claim_ any) (string, error)
	// GetClaimFromToken will return the claim from the token
	GetClaimFromToken(token_ string) (any, error) // this return any is different from the SignTokenToUser and GetUserClaimFromToken's any
	// GetUserClaimFromToken will return the client claim from the token
	GetUserClaimFromToken(token_ string) (any, error)
}

// ITokenValidator is an interface for validate each token is valid or not
type ITokenValidator interface {
	TokenValid(token_ string) bool
	TokenExpired(token_ string) bool
}

// ITokenContainer for storage token caches for it
type ITokenContainer interface {
	StoreToken(token_ string, claim_ any, duration time.Duration) error
	MarkTokenExpired(token_ string) error
	HasToken(token_ string) bool
	//HasUser(token_ string) bool
}

type IAuth interface {
	ITokenSigner
	ITokenContainer
	ITokenValidator

	MiddleWare() gin.HandlerFunc // Make a middleware for gin
}
