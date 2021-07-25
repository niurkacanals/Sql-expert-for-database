-- SQL Schema template for the a_bit_of_everything.db schema.
-- Generated on Sun Jul 25 07:09:39 WIB 2021 by xo.

CREATE TABLE a_bit_of_everything (
    a_bigint BIGINT NOT NULL,
    a_bigint_nullable BIGINT,
    a_blob BLOB NOT NULL,
    a_blob_nullable BLOB,
    a_bool BOOL NOT NULL,
    a_bool_nullable BOOL,
    a_boolean BOOLEAN NOT NULL,
    a_boolean_nullable BOOLEAN,
    a_character CHARACTER NOT NULL,
    a_character_nullable CHARACTER,
    a_clob CLOB NOT NULL,
    a_clob_nullable CLOB,
    a_date DATE NOT NULL,
    a_date_nullable DATE,
    a_datetime DATETIME NOT NULL,
    a_datetime_nullable DATETIME,
    a_decimal DECIMAL NOT NULL,
    a_decimal_nullable DECIMAL,
    a_double DOUBLE NOT NULL,
    a_double_nullable DOUBLE,
    a_float FLOAT NOT NULL,
    a_float_nullable FLOAT,
    a_int INT NOT NULL,
    a_int_nullable INT,
    a_integer INTEGER NOT NULL,
    a_integer_nullable INTEGER,
    a_mediumint MEDIUMINT NOT NULL,
    a_mediumint_nullable MEDIUMINT,
    a_native_character NATIVE CHARACTER NOT NULL,
    a_native_character_nullable NATIVE CHARACTER,
    a_nchar NCHAR NOT NULL,
    a_nchar_nullable NCHAR,
    a_numeric NUMERIC NOT NULL,
    a_numeric_nullable NUMERIC,
    a_nvarchar NVARCHAR NOT NULL,
    a_nvarchar_nullable NVARCHAR,
    a_real REAL NOT NULL,
    a_real_nullable REAL,
    a_smallint SMALLINT NOT NULL,
    a_smallint_nullable SMALLINT,
    a_text TEXT NOT NULL,
    a_text_nullable TEXT,
    a_time TIME NOT NULL,
    a_time_nullable TIME,
    a_timestamp TIMESTAMP NOT NULL,
    a_timestamp_nullable TIMESTAMP,
    a_timestamp_without_timezone TIMESTAMP WITHOUT TIMEZONE NOT NULL,
    a_timestamp_without_timezone_nullable TIMESTAMP WITHOUT TIMEZONE,
    a_timestamp_with_timezone TIMESTAMP WITH TIMEZONE NOT NULL,
    a_timestamp_with_timezone_nullable TIMESTAMP WITH TIMEZONE,
    a_time_without_timezone TIME WITHOUT TIMEZONE NOT NULL,
    a_time_without_timezone_nullable TIME WITHOUT TIMEZONE,
    a_time_with_timezone TIME WITH TIMEZONE NOT NULL,
    a_time_with_timezone_nullable TIME WITH TIMEZONE,
    a_tinyint TINYINT NOT NULL,
    a_tinyint_nullable TINYINT,
    a_varchar VARCHAR NOT NULL,
    a_varchar_nullable VARCHAR,
    a_varying_character VARYING CHARACTER NOT NULL,
    a_varying_character_nullable VARYING CHARACTER
);


CREATE TABLE a_foreign_key (
    a_key INTEGER REFERENCES a_primary (a_key)
);


CREATE TABLE a_foreign_key_composite (
    a_key1 INTEGER,
    a_key2 INTEGER,
    FOREIGN KEY (a_key1, a_key2) REFERENCES a_primary_composite (a_key1, a_key2)
);


CREATE TABLE a_index (
    a_key INTEGER
);

CREATE INDEX a_index_idx ON a_index (a_key);

CREATE TABLE a_index_composite (
    a_key1 INTEGER,
    a_key2 INTEGER
);

CREATE INDEX a_index_composite_idx ON a_index_composite (a_key1, a_key2);

CREATE TABLE a_manual_table (
    a_text VARCHAR(255)
);


CREATE TABLE a_primary (
    a_key INTEGER NOT NULL,
    PRIMARY KEY (a_key)
);


CREATE TABLE a_primary_composite (
    a_key1 INTEGER NOT NULL,
    a_key2 INTEGER NOT NULL,
    PRIMARY KEY (a_key1, a_key2)
);


CREATE TABLE a_primary_multi (
    a_key INTEGER NOT NULL,
    a_text VARCHAR(255),
    PRIMARY KEY (a_key)
);


CREATE TABLE a_sequence (
    a_seq INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT
);


CREATE TABLE a_sequence_multi (
    a_seq INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    a_text VARCHAR(255)
);


CREATE TABLE a_unique_index (
    a_key INTEGER,
    UNIQUE (a_key)
);


CREATE TABLE a_unique_index_composite (
    a_key1 INTEGER,
    a_key2 INTEGER,
    UNIQUE (a_key1, a_key2)
);


CREATE VIEW a_view_of_everything AS
  SELECT * FROM a_bit_of_everything;

CREATE VIEW a_view_of_everything_some AS
  SELECT a_bool, a_text FROM a_bit_of_everything;

