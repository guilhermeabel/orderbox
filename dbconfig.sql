-- BASE CONFIG

CREATE DATABASE IF NOT EXISTS orderbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE orderbox;

CREATE USER 'web'@'localhost' IDENTIFIED BY 'password';

GRANT ALL PRIVILEGES ON orderbox.* TO 'web'@'localhost';

FLUSH PRIVILEGES;

-- DB SCHEMA

CREATE TABLE orders (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	created DATETIME NOT NULL,
	expires DATETIME NOT NULL
);

CREATE INDEX idx_orders_created ON orders(created);


INSERT INTO orders (title, content, created, expires) VALUES (
	'An old silent pond',
	'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsuo Bashō',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO orders (title, content, created, expires) VALUES (
	'Over the wintry forest',
	'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n- Natsume Soseki',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO orders (title, content, created, expires) VALUES (
	'First autumn morning',
	'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n- Murakami Kijo',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);


-- Session table
CREATE TABLE sessions (
	token CHAR(43) PRIMARY KEY,
	data BLOB NOT NULL,
	expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

-- Users table
CREATE TABLE users (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	hashed_password CHAR(60) NOT NULL,
	created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
