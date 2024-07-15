
CREATE TABLE IF NOT EXISTS Items( 
    item_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    limits INT DEFAULT 0,
    price DECIMAL(10,2) NOT NULL,
    description TEXT NOT NULL,
    release_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expire_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    publisher BIGINT NOT NULL,
    FOREIGN KEY (publisher) REFERENCES Users(user_id)
);


CREATE TABLE IF NOT EXISTS Carts( 
    cart_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE if not exists CartedItems(
    cart_id INT NOT NULL,
    item_id INT NOT NULL,
    count INT DEFAULT 1,
    PRIMARY KEY (cart_id, item_id),
    FOREIGN KEY (cart_id) REFERENCES Carts(cart_id),
    FOREIGN KEY (item_id) REFERENCES Items(item_id)
);

CREATE TABLE IF NOT EXISTS Transactions ( 
    transaction_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
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

CREATE TABLE UserBalances (
    user_id BIGINT PRIMARY KEY,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);


CREATE INDEX idx_user_transactions ON Transactions(user_id, created_at);
CREATE INDEX idx_transaction_status ON Transactions(status);
CREATE INDEX idx_user_balances ON UserBalances(balance);
