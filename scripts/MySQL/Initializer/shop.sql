
CREATE TABLE IF NOT EXISTS Items( 
    item_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    limits INT DEFAULT 0,
    price DECIMAL(10,2) NOT NULL,
    description TEXT NOT NULL,
    release_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expire_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    publisher INT NOT NULL,
    FOREIGN KEY (publisher) REFERENCES User(user_id)
);


CREATE TABLE IF NOT EXISTS Carts( 
    cart_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES User(user_id)
);

CREATE TABLE IF NOT EXISTS Transactions ( 
    transaction_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    transaction_type ENUM('credit', 'debit') NOT NULL,
    status ENUM('pending', 'completed', 'failed') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS TransactionLogs (
    log_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    transaction_id BIGINT NOT NULL,
    log_message TEXT NOT NULL,
    logged_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES Transactions(transaction_id)
);


CREATE INDEX idx_user_transactions ON Transactions(user_id, created_at);
CREATE INDEX idx_transaction_status ON Transactions(status);
CREATE INDEX idx_user_balances ON UserBalances(balance);
