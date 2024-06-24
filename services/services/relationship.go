package services

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
	"github.com/Thenecromance/OurStories/utility/log"
	"strconv"
)

type RelationShipService interface {
	//UserHasAssociation will return a boolean value if the user has an association with other users
	UserHasAssociation(userID int, type_ int) bool

	//CreateRelationshipConnection will return a string value which it should be an url to the user's relationship connection
	CreateRelationshipConnection(userID int, type_ int) string

	// DisassociateUser will remove the user's association with other users
	DisassociateUser(userID int, associateID int) bool

	// BindingTwoUser will bind two users with the relationship
	BindingTwoUser(senderID, receiverID, relationType int)
}

type relationshipServiceImpl struct {
	repo repository.RelationshipRepository
}

func (r *relationshipServiceImpl) BindingTwoUser(senderID, receiverID, relationType int) {
	r.repo.BindTwoUser(senderID, receiverID, relationType)
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

// UserHasAssociation mean to check users has been associate relationship with others
// if user want's to bind to other friend, just don't need to care
// but if user wants to bind Couple relationship , this must be limited to 1
func (r *relationshipServiceImpl) UserHasAssociation(userID int, type_ int) bool {
	count := r.UserAssociationCount(userID, type_)
	if type_ == models.Couple {
		return count < 1 //limit the user's couple count
	}
	// temp to keep all users can have 10 friend limits
	return count < 10
}

func (r *relationshipServiceImpl) UserAssociationCount(userID int, type_ int) int {
	return r.repo.GetRelationShipCount(userID, type_)
}

// CreateRelationshipConnection a method for create bind link for users, if this user can add more friend, method will return a linkpath
// otherwise return empty string
func (r *relationshipServiceImpl) CreateRelationshipConnection(userID int, type_ int) string {
	if r.UserHasAssociation(userID, type_) {
		// todo:in this place, service will do twice sql check,this is not necessary need to be improved
		count := r.UserAssociationCount(userID, type_)
		return generateURL(userID, type_, count)
	}
	return ""
}

// DisassociateUser this method will unbind the user's relationship with other users
func (r *relationshipServiceImpl) DisassociateUser(userID int, associateID int) bool {
	if err := r.repo.UnBindTwoUser(userID, associateID); err != nil {
		log.Error("DisassociateUser failed! error: ", err, " userID: ", userID, " associateID: ", associateID)
		return false
	}
	return true
}

func NewRelationShipService(userRepository repository.RelationshipRepository) RelationShipService {
	return &relationshipServiceImpl{
		repo: userRepository,
	}
}
