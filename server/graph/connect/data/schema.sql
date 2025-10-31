-- graph/connect/data/schema.sql
-- -- DROP (dev reset)
-- DROP TABLE IF EXISTS users;

-- Users
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,   
    password VARCHAR(155) NOT NULL,
    email VARCHAR(255) NOT NULL,
    creation_date DATE NOT NULL
);