package repository

import (
	"errors"
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
	"math"
	"time"
)

type RelationshipRepository interface {
	Interface.Repository
	//HasBindToUser for checking if the user has bind to other user if they already has Bind info , return true
	//otherwise return false
	HasBindToUser(userID int, userID2 int) bool
	//GetRelationCount will return the count of the user's relationship
	//
	//	userID int - the user's id
	//
	//	type_ int - the relation type
	//
	//	return int - the count of the user's relationship
	GetRelationCount(userID int, type_ int) int

	//BindTwoUser will bind two users with the relationship
	//
	//	senderID int - the sender's id
	//
	//	receiverID int - the receiver's id
	//
	//	relationType int - the relation type
	//
	//	return error - if the operation failed, return the error
	BindTwoUser(senderID, receiverID, relationType int) error

	//UnBindTwoUser will unbind two users with the relationship
	//
	//	senderID int - the sender's id
	//
	//	receiverID int - the receiver's id
	//
	// return error - if the operation failed, return the error
	UnBindTwoUser(senderID, receiverID int) error

	GetRelationList(userID int) []models.Relationship

	GetHistoryList(userID int) []models.RelationShipHistory
}

type relationshipRepositoryImpl struct {
	db *gorp.DbMap
}

func (r *relationshipRepositoryImpl) HasBindToUser(userID int, userID2 int) bool {
	count, err := r.db.SelectInt("SELECT COUNT(*) FROM relationship WHERE (user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, userID2, userID2, userID)

	if err != nil {
		log.Error("HasBindToUser failed! error: ", err, " userID: ", userID, " userID2: ", userID2)
		return false
	}

	if count > 0 {
		return true
	}

	return false
}

func (r *relationshipRepositoryImpl) GetRelationCount(userID int, type_ int) int {
	selectInt, err := r.db.SelectInt("SELECT COUNT(*) FROM relationship WHERE user_id = ? AND relation_type = ?", userID, type_)
	if err != nil {
		log.Error("GetRelationCount failed! error: ", err, " userID: ", userID, " type: ", type_)
		return math.MinInt
	}
	return int(selectInt)
}

func (r *relationshipRepositoryImpl) BindTwoUser(senderID, receiverID, relationType int) error {
	now := time.Now().Unix()

	obj, err := r.db.Select(models.Relationship{}, "SELECT * FROM relationship WHERE (user_id = ? AND friend_id = ? ) OR (user_id = ? AND friend_id = ?)", senderID, receiverID, receiverID, senderID)
	if err != nil || (obj != nil && len(obj) > 0) {
		return errors.New("the relationship already exists")
	}

	transaction, err := r.db.Begin()
	if err != nil {
		log.Error("Transaction create failed! error: ", err, " senderID: ", senderID, " receiverID: ", receiverID, " relationType: ", relationType)
		return err
	}

	//record Operation time
	{
		err := r.recordOperationTime(transaction,
			senderID,
			receiverID,
			relationType,
			models.Binding,
			now)
		if err != nil {
			err = transaction.Rollback()
			return err
		}
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

func (r *relationshipRepositoryImpl) recordOperationTime(transaction *gorp.Transaction, senderID, receiverID, relationType int, operationType int, timestamp int64) error {
	history := models.RelationShipHistory{
		UserID:        senderID,
		ReceiverID:    receiverID,
		OperationType: relationType,
		OperationTime: timestamp,
		Operation:     operationType,
		OperationUser: senderID, // this is unique
	}
	err := transaction.Insert(&history)
	if err != nil {
		return err
	}

	/*	history.UserID = receiverID
		history.ReceiverID = senderID
		err = transaction.Insert(&history)*/
	return err
}

func (r *relationshipRepositoryImpl) UnBindTwoUser(senderID, receiverID int) error {
	transaction, err := r.db.Begin()
	if err != nil {
		log.Error("Transaction create failed! error: ", err, " senderID: ", senderID, " receiverID: ", receiverID)
		return err
	}

	//record Operation time
	{
		r.recordOperationTime(transaction, senderID, receiverID, models.Unknown, models.Disassociate, time.Now().Unix())
		if err != nil {
			err = transaction.Rollback()
			if err != nil {
				return err
			}
			return err
		}
	}

	_, err = transaction.Exec("DELETE FROM relationship WHERE (user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", senderID, receiverID, receiverID, senderID)
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

func (r *relationshipRepositoryImpl) GetRelationList(userID int) []models.Relationship {
	result, err := r.db.Select(models.Relationship{}, "SELECT * FROM relationship WHERE user_id = ? OR friend_id = ?", userID, userID)
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

func (r *relationshipRepositoryImpl) GetHistoryList(userID int) []models.RelationShipHistory {
	history, err := r.db.Select(models.RelationShipHistory{}, "SELECT * FROM relationship_history WHERE user_id = ? OR target_id = ?", userID, userID)
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

/*func (r *relationshipRepositoryImpl) initTable() error {
	if r.db == nil {
		log.Debugf("db is nil")
		return fmt.Errorf("db is nil")
	}

	{
		tableName := "relationship"
		log.Infof("start to binding %s table", tableName)
		tbl := r.db.AddTableWithName(models.Relationship{}, tableName)
		tbl.SetKeys(true, "id")
	}
	{
		tableName := "relationship_history"
		log.Infof("start to binding user %s table", tableName)
		tbl := r.db.AddTableWithName(models.RelationShipHistory{}, tableName)
		tbl.SetKeys(true, "id")
	}

	return r.db.CreateTablesIfNotExists()
}*/

func (r *relationshipRepositoryImpl) BindTable() error {
	r.db.AddTableWithName(models.Relationship{}, "Relations")
	r.db.AddTableWithName(models.RelationShipHistory{}, "RelationLogs")
	return nil
}

func NewRelationShipRepository(db *gorp.DbMap) RelationshipRepository {
	repo := &relationshipRepositoryImpl{
		db: db,
	}
	//repo.initTable()
	return repo
}
