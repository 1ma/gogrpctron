SET NAMES 'utf8mb4';

CREATE TABLE users (
    id         INT          NOT NULL AUTO_INCREMENT,
    username   VARCHAR(255) NOT NULL,
    hash       VARCHAR(60)  NOT NULL,
    created_at DATETIME     NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX idx_users_created_at ON users(created_at);
