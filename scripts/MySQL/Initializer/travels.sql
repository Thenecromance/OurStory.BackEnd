
CREATE TABLE IF NOT EXISTS Travels ( 
    travel_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    state INT NOT NULL , 
    travel_date BIGINT NOT NULL,
    travel_time BIGINT NOT NULL,
    travel_from VARCHAR(255) NOT NULL,
    travel_to VARCHAR(255) NOT NULL,
    travel_cost DECIMAL(10,2) NOT NULL,
    created_at BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS TravelLogs(
    log_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    travel_id BIGINT NOT NULL,
    modified_by BIGINT NOT NULL,
    message TEXT NOT NULL,
    modified_at BIGINT NOT NULL,
    FOREIGN KEY (travel_id) REFERENCES Travels(travel_id),
    FOREIGN KEY (modified_by) REFERENCES Users(user_id)
);