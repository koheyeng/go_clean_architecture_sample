CREATE TABLE IF NOT EXISTS edge_enable (
    id INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(36) NOT NULL,
    age INTEGER NOT NULL,
    birthday TIMESTAMP,
    address VARCHAR(128),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)