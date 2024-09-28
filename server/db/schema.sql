CREATE TABLE users (
    id serial PRIMARY KEY,
    firebase_uid varchar(255) NOT NULL UNIQUE,
    username varchar(50) NOT NULL,
    email varchar(100) NOT NULL UNIQUE,
    password varchar(255) NOT NULL
);

CREATE TABLE todos (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id),
    text varchar(255) NOT NULL,
    status varchar(20) NOT NULL DEFAULT 'pending',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)