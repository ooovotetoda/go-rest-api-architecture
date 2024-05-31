CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     phone VARCHAR(15) UNIQUE NOT NULL,
                                     name VARCHAR(255) NOT NULL,
                                     age INT NOT NULL
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          description TEXT,
                          price DECIMAL(10, 2) NOT NULL
);
