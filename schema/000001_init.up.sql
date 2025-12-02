CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE todo_lists (
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text
);

CREATE TABLE users_lists (
    id serial PRIMARY KEY,
    user_id int REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    list_id int REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL,
    UNIQUE(user_id, list_id)
);

CREATE TABLE todo_items (
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text,
    done boolean NOT NULL DEFAULT false
);

CREATE TABLE lists_items (
    id serial PRIMARY KEY,
    item_id int REFERENCES todo_items(id) ON DELETE CASCADE NOT NULL,
    list_id int REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL,
    UNIQUE(item_id, list_id)
);