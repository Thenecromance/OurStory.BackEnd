package services

import (
	"errors"
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/application/services/RelationValidator"
	"github.com/Thenecromance/OurStories/utility/log"
)

type RelationShipService interface {
	//UserHasAssociation will return a boolean value if the user has an association with other users
	UserHasAssociation(userID int64, type_ int) bool

	//CreateRelationshipConnection will return a string value which it should be an url to the user's relationship connection
	CreateRelationshipConnection(userID int64, type_ int) string

	// DisassociateUser will remove the user's association with other users
	DisassociateUser(userID int64, associateID int64) bool

	// BindingTwoUser will bind two users with the relationship
	BindingTwoUser(link string, receiverID int64) (err error)

	GetFriendList(userID int64) []models.Relationship

	GetHistoryList(userID int64) []models.RelationShipHistory
}

type relationshipServiceImpl struct {
	repo      repository.RelationshipRepository
	userRepo  repository.UserRepository
	validator RelationValidator.IRelationValidator
}

func (r *relationshipServiceImpl) GetFriendList(userID int64) []models.Relationship {
	lists := r.repo.GetRelationList(userID)

	return lists
}

func (r *relationshipServiceImpl) BindingTwoUser(link string, receiverID int64) error {

	senderID, relationType, err := r.validator.GetTokenInfo(link)
	if err != nil {
		log.Warn("BindingTwoUser failed! error: ", err)
		return err
	}

	if r.userRepo.HasId(receiverID) == false {
		log.Warn("BindingTwoUser failed! error: ", "receiverID not exists")
		return errors.New("receiverID not exists")
	}

	//in case to avoid the senderID equals receiverID
	if senderID == receiverID {
		log.Warn("BindingTwoUser failed! error: ", "senderID equals receiverID")
		return errors.New("senderID equals receiverID")
	}

	err = r.repo.BindTwoUser(senderID, receiverID, relationType)
	if err != nil {
		log.Warn("BindingTwoUser failed! error: ", err)
		return err
	}
	return err
}

// UserHasAssociation mean to check users has been associate relationship with others
// if user want's to bind to other friend, just don't need to care
// but if user wants to bind Couple relationship , this must be limited to 1
func (r *relationshipServiceImpl) UserHasAssociation(userID int64, type_ int) bool {
	count := r.UserAssociationCount(userID, type_)

	if type_ == models.Couple {
		return count < 1 //limit the user's couple count
	}
	// temp to keep all users can have 10 friend limits
	return count < 10
}

func (r *relationshipServiceImpl) UserAssociationCount(userID int64, type_ int) int {
	return r.repo.GetRelationCount(userID, type_)
}

// CreateRelationshipConnection a method for create bind link for users, if this user can add more friend, method will return a linkpath
// otherwise return empty string
func (r *relationshipServiceImpl) CreateRelationshipConnection(userID int64, type_ int) string {
	if r.UserHasAssociation(userID, type_) {
		// todo:in this place, service will do twice sql check,this is not necessary need to be improved
		count := r.UserAssociationCount(userID, type_)
		token, err := r.validator.GenerateToken(userID, type_, count)

		if err != nil {
			log.Error(err)
			return ""
		}
		return token
	}
	return ""
}

// DisassociateUser this method will unbind the user's relationship with other users
func (r *relationshipServiceImpl) DisassociateUser(userID int64, associateID int64) bool {
	if err := r.repo.UnBindTwoUser(userID, associateID); err != nil {
		log.Error("DisassociateUser failed! error: ", err, " userID: ", userID, " associateID: ", associateID)
		return false
	}
	return true
}

func (r *relationshipServiceImpl) GetHistoryList(userID int64) []models.RelationShipHistory {
	return r.repo.GetHistoryList(userID)
}

func NewRelationShipService(userRepository repository.RelationshipRepository) RelationShipService {
	return &relationshipServiceImpl{
		repo:      userRepository,
		validator: RelationValidator.New(),
	}
}
