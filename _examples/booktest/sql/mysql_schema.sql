SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS books;
DROP FUNCTION IF EXISTS say_hello;
SET FOREIGN_KEY_CHECKS=1;

CREATE TABLE authors (
  author_id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name TEXT NOT NULL DEFAULT ''
) ENGINE=InnoDB;

CREATE INDEX authors_name_idx ON authors(name(255));

CREATE TABLE books (
  book_id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
  author_id INTEGER NOT NULL REFERENCES authors(author_id),
  isbn VARCHAR(255) NOT NULL DEFAULT '' UNIQUE,
  book_type ENUM('FICTION', 'NONFICTION') NOT NULL DEFAULT 'FICTION',
  title TEXT NOT NULL DEFAULT '',
  year INTEGER NOT NULL DEFAULT 2000,
  available DATETIME NOT NULL DEFAULT NOW(),
  tags TEXT NOT NULL DEFAULT ''
) ENGINE=InnoDB;

CREATE INDEX books_title_idx ON books(title(255), year);

CREATE FUNCTION say_hello(s TEXT) RETURNS TEXT
  DETERMINISTIC
  RETURN CONCAT('hello ', s);
