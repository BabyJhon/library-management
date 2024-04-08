CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    sure_name varchar(255) NOT NULL,
    phone_number varchar(255) NOT NULL
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title varchar(255) NOT NULL,
    author varchar(255) NOT NULL,
    in_library boolean
);

CREATE TABLE users_books (
    id SERIAL PRIMARY KEY,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    book_id int,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    sure_name varchar(255) NOT NULL,
    user_name varchar(255) NOT NULL,
    password_hash varchar(255) NOT NULL
)