CREATE TABLE IF NOT EXISTS Anniversaries (
    anniversary_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL, 
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    anniversary_date BIGINT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL, 
    shared_with TEXT NOT NULL, 
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);