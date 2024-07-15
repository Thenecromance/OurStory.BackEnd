
CREATE TABLE IF NOT EXISTS Travels ( 
    travel_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    travel_date DATE NOT NULL,
    travel_time TIME NOT NULL,
    travel_from VARCHAR(255) NOT NULL,
    travel_to VARCHAR(255) NOT NULL,
    travel_cost DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE if not exists TravelLogs(
    log_id INT PRIMARY KEY AUTO_INCREMENT,
    travel_id INT NOT NULL,
    modified_by BIGINT NOT NULL,
    log_message TEXT NOT NULL,
    logged_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (travel_id) REFERENCES Travels(travel_id),
    FOREIGN KEY (modified_by) REFERENCES Users(user_id)
);