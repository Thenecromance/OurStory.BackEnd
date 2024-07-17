
CREATE TABLE IF NOT EXISTS Travels ( 
    travel_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    state INT NOT NULL , 
    travel_start BIGINT NOT NULL,
    travel_end BIGINT NOT NULL,
    location VARCHAR(255) NOT NULL,
    detail TEXT NOT NULL,
    together TEXT NOT NULL, 
    image VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
    on update cascade 
    on delete cascade
);

CREATE TABLE IF NOT EXISTS TravelLogs(
    log_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    travel_id BIGINT NOT NULL,
    modified_by BIGINT NOT NULL,
    message TEXT NOT NULL,
    modified_at BIGINT NOT NULL,
    FOREIGN KEY (travel_id) REFERENCES Travels(travel_id),
    FOREIGN KEY (modified_by) REFERENCES Users(user_id)
    on update cascade 
    on delete cascade
);