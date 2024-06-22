package services

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/Thenecromance/OurStories/services/repository"
	"strconv"
)

type RelationShipService interface {
	//UserHasAssociation will return a boolean value if the user has an association with other users
	UserHasAssociation(userID int) bool
	//CreateRelationshipConnection will return a string value which it should be an url to the user's relationship connection
	CreateRelationshipConnection(userID int) string
	// DisassociateUser will remove the user's association with other users
	DisassociateUser(userID int) bool
}

type relationshipServiceImpl struct {
	repo repository.UserRepository
}

// generateURL
// generate a unique url based on the user's id and the relation type
//
//	userID int - the user's id
//	relationType int - the relation type
//	idx int - the index of the relations, exp: if the user has more than 1 friend, the idx will be 1, 2, 3
func generateURL(userID, relationType, idx int) string {
	data := strconv.Itoa(userID) + "." + strconv.Itoa(relationType) + "." + strconv.Itoa(idx)
	hash := sha256.Sum256([]byte(data))
	uniqueID := hex.EncodeToString(hash[:])
	return base64.StdEncoding.EncodeToString([]byte(uniqueID)) // the link should like https://.../api/relation/bind/NzdhYzMxOWJmZTE5NzllMmQ3OTlkOWU2OTg3ZTY1ZmViNTRmNjE1MTFjMDM1NTJlYmFlOTkwODI2YzIwODU5MA==
}

func (r *relationshipServiceImpl) UserHasAssociation(userID int) bool {
	//TODO implement me
	panic("implement me")
}

func (r *relationshipServiceImpl) CreateRelationshipConnection(userID int) string {
	//TODO implement me
	panic("implement me")
}

func (r *relationshipServiceImpl) DisassociateUser(userID int) bool {
	//TODO implement me
	panic("implement me")
}

func NewRelationShipService(userRepository repository.UserRepository) RelationShipService {
	return &relationshipServiceImpl{
		repo: userRepository,
	}
}
