package RelationValidator

import "errors"

type IRelationValidator interface {
	// GenerateToken will generate a token based on the given userID, relationType and idx
	GenerateToken(userID int64, relationType int, idx int) (string, error)
	// GetTokenInfo will return all details of the token if it exists
	GetTokenInfo(token string) (userID int64, relationType int, err error)
}

var (
	ErrTokenNotFount = errors.New("no token found in cache")
)
