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
