
/*
type Relationship struct {
	ID            int    `json:"id" db:"id"`
	UserID        int    `json:"user_id" db:"user_id" `            // the user id
	FriendID      int    `json:"friend_id" db:"friend_id"`         // associate with the user id
	RelationType  int    `json:"relation_type" db:"relation_type"` // two of the user's relationship type
	Status        string `json:"status" db:"status"`               // the status of the relationship
	AssociateTime int64  `json:"stamp" db:"associate_time"`        // the time when the relationship is created
}
*/


CREATE TABLE if not exists Relations (
    relation_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    friend_id BIGINT NOT NULL,
    relation_type INT NOT NULL,
    status ENUM('pending', 'accepted', 'rejected') NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (friend_id) REFERENCES Users(user_id)
);

CREATE TABLE if not exists RelationLogs(
    log_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,  
    target_id BIGINT NOT NULL ,  
    operation_user BIGINT NOT NULL ,
    operation_type INT NOT NULL, 
    operation_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    operation INT NOT NULL , 
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (target_id) REFERENCES Users(user_id),
    FOREIGN KEY(operation_user) REFERENCES Users(user_id)
);