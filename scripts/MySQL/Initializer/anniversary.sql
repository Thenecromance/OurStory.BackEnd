CREATE TABLE IF not exists Anniversaries (
    anniversary_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL, 
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    anniversary_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);