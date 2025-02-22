package repository

import (
	"errors"
	"math"
	"time"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

type RelationshipRepository interface {
	Interface.Repository
	//HasBindToUser for checking if the user has bind to other user if they already has Bind info , return true
	//otherwise return false
	HasBindToUser(userID int64, userID2 int64) bool
	//GetRelationCount will return the count of the user's relationship
	//
	//	userID int - the user's id
	//
	//	type_ int - the relation type
	//
	//	return int - the count of the user's relationship
	GetRelationCount(userID int64, type_ int) int

	//BindTwoUser will bind two users with the relationship
	//
	//	senderID int - the sender's id
	//
	//	receiverID int - the receiver's id
	//
	//	relationType int - the relation type
	//
	//	return error - if the operation failed, return the error
	BindTwoUser(senderID, receiverID int64, relationType int) error

	//UnBindTwoUser will unbind two users with the relationship
	//
	//	senderID int - the sender's id
	//
	//	receiverID int - the receiver's id
	//
	// return error - if the operation failed, return the error
	UnBindTwoUser(senderID, receiverID int64) error

	GetRelationList(userID int64) []models.Relationship

	GetHistoryList(userID int64) []models.RelationShipHistory
}

type relationshipRepositoryImpl struct {
	db *gorp.DbMap
}

func (r *relationshipRepositoryImpl) HasBindToUser(userID int64, userID2 int64) bool {

	count, err := r.db.SelectInt("SELECT COUNT(*) FROM Relations WHERE (user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, userID2, userID2, userID)

	if err != nil {
		log.Error("HasBindToUser failed! error: ", err, " userID: ", userID, " userID2: ", userID2)
		return false
	}

	if count > 0 {
		return true
	}

	return false
}

func (r *relationshipRepositoryImpl) GetRelationCount(userID int64, type_ int) int {
	selectInt, err := r.db.SelectInt("SELECT COUNT(*) FROM Relations WHERE user_id = ? AND relation_type = ?", userID, type_)
	if err != nil {
		log.Error("GetRelationCount failed! error: ", err, " userID: ", userID, " type: ", type_)
		return math.MinInt
	}
	return int(selectInt)
}

func (r *relationshipRepositoryImpl) BindTwoUser(senderID, receiverID int64, relationType int) error {
	now := time.Now().Unix()

	obj, err := r.db.Select(models.Relationship{}, "SELECT * FROM Relations WHERE (user_id = ? AND friend_id = ? ) OR (user_id = ? AND friend_id = ?)", senderID, receiverID, receiverID, senderID)
	if err != nil || (obj != nil && len(obj) > 0) {
		return errors.New("the relationship already exists")
	}

	transaction, err := r.db.Begin()
	if err != nil {
		log.Error("Transaction create failed! error: ", err, " senderID: ", senderID, " receiverID: ", receiverID, " relationType: ", relationType)
		return err
	}

	//record relationship info
	{
		relationship := models.Relationship{
			UserID:        senderID,
			FriendID:      receiverID,
			RelationType:  relationType,
			AssociateTime: now, // maybe this part not necessary
		}
		err = transaction.Insert(&relationship)
		if err != nil {
			err = transaction.Rollback()
			if err != nil {
				return err
			}
		}
	}

	err = transaction.Commit()
	if err != nil {
		log.Error("Transaction commit failed! error: ", err, " senderID: ", senderID, " receiverID: ", receiverID, " relationType: ", relationType)
		transaction.Rollback()
		return err
	}
	return nil
}

func (r *relationshipRepositoryImpl) UnBindTwoUser(senderID, receiverID int64) error {
	transaction, err := r.db.Begin()
	if err != nil {
		log.Error("Transaction create failed! error: ", err, " senderID: ", senderID, " receiverID: ", receiverID)
		return err
	}

	_, err = transaction.Exec("DELETE FROM Relations WHERE (user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", senderID, receiverID, receiverID, senderID)
	if err != nil {
		err = transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = transaction.Commit()
	if err != nil {
		log.Error("Transaction commit failed! error: ", err, " senderID: ", senderID, " receiverID: ", receiverID)
		transaction.Rollback()
		return err
	}
	return nil

}

func (r *relationshipRepositoryImpl) GetRelationList(userID int64) []models.Relationship {
	result, err := r.db.Select(models.Relationship{}, "SELECT * FROM Relations WHERE user_id = ? OR friend_id = ?", userID, userID)
	if err != nil {
		log.Warn("GetRelationList failed! error: ", err, " userID: ", userID)
		return nil
	}
	var relationships []models.Relationship
	for _, item := range result {
		relationships = append(relationships, *item.(*models.Relationship))
	}
	return relationships
}

func (r *relationshipRepositoryImpl) GetHistoryList(userID int64) []models.RelationShipHistory {
	history, err := r.db.Select(models.RelationShipHistory{}, "SELECT * FROM RelationLogs WHERE user_id = ? OR target_id = ?", userID, userID)
	if err != nil {
		log.Warn("GetHistoryList failed! error: ", err, " userID: ", userID)
		return nil
	}
	//format []interface{} to []models.RelationShipHistory
	var histories []models.RelationShipHistory
	for _, item := range history {
		histories = append(histories, *item.(*models.RelationShipHistory))
	}
	return histories
}

func (r *relationshipRepositoryImpl) BindTable() error {
	r.db.AddTableWithName(models.Relationship{}, "Relations")
	r.db.AddTableWithName(models.RelationShipHistory{}, "RelationLogs")
	return nil
}

func NewRelationShipRepository(db *gorp.DbMap) RelationshipRepository {
	repo := &relationshipRepositoryImpl{
		db: db,
	}

	return repo
}
