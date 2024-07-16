CREATE table IF NOT EXISTS Users (
    user_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) DEFAULT 'default.png',
    nickname VARCHAR(255) DEFAULT NULL,
    role INT DEFAULT 0 , 
    email VARCHAR(255) NOT NULL,
    birthday BIGINT NOT NULL,
    gender VARCHAR(10) DEFAULT 'unknown',
    created_at  BIGINT NOT NULL,
    last_login  BIGINT NOT NULL,
    pass_word VARCHAR(255) NOT NULL ,
    salt VARCHAR(255) NOT NULL
);



CREATE TABLE IF NOT EXISTS LoginLogs(
    user_id BIGINT NOT NULL , 
    login_time BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);