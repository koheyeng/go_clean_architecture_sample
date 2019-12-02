CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY NOT NULL,
    first_name VARCHAR(36) NOT NULL,
    family_name VARCHAR(36) NOT NULL,
    age INTEGER NOT NULL,
    birthday TIMESTAMP,
    address VARCHAR(128),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)