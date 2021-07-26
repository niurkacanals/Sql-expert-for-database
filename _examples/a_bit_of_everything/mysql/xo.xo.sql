-- Generated by xo for the a_bit_of_everything schema.

-- table a_bit_of_everything
CREATE TABLE a_bit_of_everything (
  a_bigint BIGINT(20) NOT NULL,
  a_bigint_nullable BIGINT(20),
  a_binary BINARY(1) NOT NULL,
  a_binary_nullable BINARY(1),
  a_bit BIT(1) NOT NULL,
  a_bit_nullable BIT(1),
  a_blob BLOB NOT NULL,
  a_blob_nullable BLOB,
  a_bool TINYINT(1) NOT NULL,
  a_bool_nullable TINYINT(1),
  a_char CHAR(1) NOT NULL,
  a_char_nullable CHAR(1),
  a_date DATE NOT NULL,
  a_date_nullable DATE,
  a_datetime DATETIME NOT NULL,
  a_datetime_nullable DATETIME,
  a_dec DECIMAL(10) NOT NULL,
  a_dec_nullable DECIMAL(10),
  a_fixed DECIMAL(10) NOT NULL,
  a_fixed_nullable DECIMAL(10),
  a_decimal DECIMAL(10) NOT NULL,
  a_decimal_nullable DECIMAL(10),
  a_double_precision DOUBLE NOT NULL,
  a_double_precision_nullable DOUBLE,
  a_enum ENUM('ONE', 'TWO') NOT NULL,
  a_enum_nullable ENUM('ONE', 'TWO'),
  a_float FLOAT NOT NULL,
  a_float_nullable FLOAT,
  a_int INT(11) NOT NULL,
  a_int_nullable INT(11),
  a_integer INT(11) NOT NULL,
  a_integer_nullable INT(11),
  a_json LONGTEXT NOT NULL,
  a_json_nullable LONGTEXT,
  a_longblob LONGBLOB NOT NULL,
  a_longblob_nullable LONGBLOB,
  a_longtext LONGTEXT NOT NULL,
  a_longtext_nullable LONGTEXT,
  a_mediumblob MEDIUMBLOB NOT NULL,
  a_mediumblob_nullable MEDIUMBLOB,
  a_mediumint MEDIUMINT(9) NOT NULL,
  a_mediumint_nullable MEDIUMINT(9),
  a_mediumtext MEDIUMTEXT NOT NULL,
  a_mediumtext_nullable MEDIUMTEXT,
  a_numeric DECIMAL(10) NOT NULL,
  a_numeric_nullable DECIMAL(10),
  a_real FLOAT NOT NULL,
  a_real_nullable FLOAT,
  a_set SET('ONE','TWO') NOT NULL,
  a_set_nullable SET('ONE','TWO'),
  a_smallint SMALLINT(6) NOT NULL,
  a_smallint_nullable SMALLINT(6),
  a_text TEXT NOT NULL,
  a_text_nullable TEXT,
  a_time TIME NOT NULL,
  a_time_nullable TIME,
  a_timestamp TIMESTAMP DEFAULT current_timestamp() NOT NULL,
  a_timestamp_nullable TIMESTAMP DEFAULT '0000-00-00 00:00:00' NOT NULL,
  a_tinyblob TINYBLOB NOT NULL,
  a_tinyblob_nullable TINYBLOB,
  a_tinyint TINYINT(4) NOT NULL,
  a_tinyint_nullable TINYINT(4),
  a_tinytext TINYTEXT NOT NULL,
  a_tinytext_nullable TINYTEXT,
  a_varbinary VARBINARY(255) NOT NULL,
  a_varbinary_nullable VARBINARY(255),
  a_varchar VARCHAR(255) NOT NULL,
  a_varchar_nullable VARCHAR(255),
  a_year YEAR(4) NOT NULL,
  a_year_nullable YEAR(4)
) ENGINE=InnoDB;

-- table a_primary
CREATE TABLE a_primary (
  a_key INT(11) NOT NULL,
  PRIMARY KEY (a_key)
) ENGINE=InnoDB;

-- table a_foreign_key
CREATE TABLE a_foreign_key (
  a_key INT(11) REFERENCES a_primary (a_key)
) ENGINE=InnoDB;

-- index a_key
CREATE INDEX a_key ON a_foreign_key (a_key);

-- table a_primary_composite
CREATE TABLE a_primary_composite (
  a_key1 INT(11) NOT NULL,
  a_key2 INT(11) NOT NULL,
  PRIMARY KEY (a_key1, a_key2)
) ENGINE=InnoDB;

-- table a_foreign_key_composite
CREATE TABLE a_foreign_key_composite (
  a_key1 INT(11),
  a_key2 INT(11),
  FOREIGN KEY (a_key1, a_key2) REFERENCES a_primary_composite (a_key1, a_key2)
) ENGINE=InnoDB;

-- index a_key1
CREATE INDEX a_key1 ON a_foreign_key_composite (a_key1, a_key2);

-- table a_index
CREATE TABLE a_index (
  a_key INT(11)
) ENGINE=InnoDB;

-- index a_index_idx
CREATE INDEX a_index_idx ON a_index (a_key);

-- table a_index_composite
CREATE TABLE a_index_composite (
  a_key1 INT(11),
  a_key2 INT(11)
) ENGINE=InnoDB;

-- index a_index_composite_idx
CREATE INDEX a_index_composite_idx ON a_index_composite (a_key1, a_key2);

-- table a_manual_table
CREATE TABLE a_manual_table (
  a_text VARCHAR(255)
) ENGINE=InnoDB;

-- table a_primary_multi
CREATE TABLE a_primary_multi (
  a_key INT(11) NOT NULL,
  a_text VARCHAR(255),
  PRIMARY KEY (a_key)
) ENGINE=InnoDB;

-- table a_sequence
CREATE TABLE a_sequence (
  a_seq INT(11) AUTO_INCREMENT,
  PRIMARY KEY (a_seq)
) ENGINE=InnoDB;

-- table a_sequence_multi
CREATE TABLE a_sequence_multi (
  a_seq INT(11) AUTO_INCREMENT,
  a_text VARCHAR(255),
  PRIMARY KEY (a_seq)
) ENGINE=InnoDB;

-- table a_unique_index
CREATE TABLE a_unique_index (
  a_key INT(11),
  UNIQUE (a_key)
) ENGINE=InnoDB;

-- table a_unique_index_composite
CREATE TABLE a_unique_index_composite (
  a_key1 INT(11),
  a_key2 INT(11),
  UNIQUE (a_key1, a_key2)
) ENGINE=InnoDB;

-- view a_view_of_everything
CREATE VIEW a_view_of_everything AS
select `a_bit_of_everything`.`a_bit_of_everything`.`a_bigint` AS `a_bigint`,`a_bit_of_everything`.`a_bit_of_everything`.`a_bigint_nullable` AS `a_bigint_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_binary` AS `a_binary`,`a_bit_of_everything`.`a_bit_of_everything`.`a_binary_nullable` AS `a_binary_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_bit` AS `a_bit`,`a_bit_of_everything`.`a_bit_of_everything`.`a_bit_nullable` AS `a_bit_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_blob` AS `a_blob`,`a_bit_of_everything`.`a_bit_of_everything`.`a_blob_nullable` AS `a_blob_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_bool` AS `a_bool`,`a_bit_of_everything`.`a_bit_of_everything`.`a_bool_nullable` AS `a_bool_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_char` AS `a_char`,`a_bit_of_everything`.`a_bit_of_everything`.`a_char_nullable` AS `a_char_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_date` AS `a_date`,`a_bit_of_everything`.`a_bit_of_everything`.`a_date_nullable` AS `a_date_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_datetime` AS `a_datetime`,`a_bit_of_everything`.`a_bit_of_everything`.`a_datetime_nullable` AS `a_datetime_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_dec` AS `a_dec`,`a_bit_of_everything`.`a_bit_of_everything`.`a_dec_nullable` AS `a_dec_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_fixed` AS `a_fixed`,`a_bit_of_everything`.`a_bit_of_everything`.`a_fixed_nullable` AS `a_fixed_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_decimal` AS `a_decimal`,`a_bit_of_everything`.`a_bit_of_everything`.`a_decimal_nullable` AS `a_decimal_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_double_precision` AS `a_double_precision`,`a_bit_of_everything`.`a_bit_of_everything`.`a_double_precision_nullable` AS `a_double_precision_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_enum` AS `a_enum`,`a_bit_of_everything`.`a_bit_of_everything`.`a_enum_nullable` AS `a_enum_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_float` AS `a_float`,`a_bit_of_everything`.`a_bit_of_everything`.`a_float_nullable` AS `a_float_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_int` AS `a_int`,`a_bit_of_everything`.`a_bit_of_everything`.`a_int_nullable` AS `a_int_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_integer` AS `a_integer`,`a_bit_of_everything`.`a_bit_of_everything`.`a_integer_nullable` AS `a_integer_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_json` AS `a_json`,`a_bit_of_everything`.`a_bit_of_everything`.`a_json_nullable` AS `a_json_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_longblob` AS `a_longblob`,`a_bit_of_everything`.`a_bit_of_everything`.`a_longblob_nullable` AS `a_longblob_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_longtext` AS `a_longtext`,`a_bit_of_everything`.`a_bit_of_everything`.`a_longtext_nullable` AS `a_longtext_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_mediumblob` AS `a_mediumblob`,`a_bit_of_everything`.`a_bit_of_everything`.`a_mediumblob_nullable` AS `a_mediumblob_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_mediumint` AS `a_mediumint`,`a_bit_of_everything`.`a_bit_of_everything`.`a_mediumint_nullable` AS `a_mediumint_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_mediumtext` AS `a_mediumtext`,`a_bit_of_everything`.`a_bit_of_everything`.`a_mediumtext_nullable` AS `a_mediumtext_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_numeric` AS `a_numeric`,`a_bit_of_everything`.`a_bit_of_everything`.`a_numeric_nullable` AS `a_numeric_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_real` AS `a_real`,`a_bit_of_everything`.`a_bit_of_everything`.`a_real_nullable` AS `a_real_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_set` AS `a_set`,`a_bit_of_everything`.`a_bit_of_everything`.`a_set_nullable` AS `a_set_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_smallint` AS `a_smallint`,`a_bit_of_everything`.`a_bit_of_everything`.`a_smallint_nullable` AS `a_smallint_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_text` AS `a_text`,`a_bit_of_everything`.`a_bit_of_everything`.`a_text_nullable` AS `a_text_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_time` AS `a_time`,`a_bit_of_everything`.`a_bit_of_everything`.`a_time_nullable` AS `a_time_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_timestamp` AS `a_timestamp`,`a_bit_of_everything`.`a_bit_of_everything`.`a_timestamp_nullable` AS `a_timestamp_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_tinyblob` AS `a_tinyblob`,`a_bit_of_everything`.`a_bit_of_everything`.`a_tinyblob_nullable` AS `a_tinyblob_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_tinyint` AS `a_tinyint`,`a_bit_of_everything`.`a_bit_of_everything`.`a_tinyint_nullable` AS `a_tinyint_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_tinytext` AS `a_tinytext`,`a_bit_of_everything`.`a_bit_of_everything`.`a_tinytext_nullable` AS `a_tinytext_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_varbinary` AS `a_varbinary`,`a_bit_of_everything`.`a_bit_of_everything`.`a_varbinary_nullable` AS `a_varbinary_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_varchar` AS `a_varchar`,`a_bit_of_everything`.`a_bit_of_everything`.`a_varchar_nullable` AS `a_varchar_nullable`,`a_bit_of_everything`.`a_bit_of_everything`.`a_year` AS `a_year`,`a_bit_of_everything`.`a_bit_of_everything`.`a_year_nullable` AS `a_year_nullable` from `a_bit_of_everything`.`a_bit_of_everything`;

-- view a_view_of_everything_some
CREATE VIEW a_view_of_everything_some AS
select `a_bit_of_everything`.`a_bit_of_everything`.`a_bool` AS `a_bool`,`a_bit_of_everything`.`a_bit_of_everything`.`a_text` AS `a_text` from `a_bit_of_everything`.`a_bit_of_everything`;

-- procedure a_0_in_0_out
CREATE PROCEDURE a_0_in_0_out()
BEGIN
END;

-- procedure a_0_in_1_out
CREATE PROCEDURE a_0_in_1_out() RETURNS INT(11)
BEGIN
  SELECT 10 INTO a_return\;
END;

-- procedure a_1_in_0_out
CREATE PROCEDURE a_1_in_0_out(a_param INT(11), OUT a_param INT(11))
BEGIN
END;

-- procedure a_1_in_1_out
CREATE PROCEDURE a_1_in_1_out(a_param INT(11)) RETURNS INT(11)
BEGIN
  SELECT a_param INTO a_return\;
END;

-- procedure a_2_in_2_out
CREATE PROCEDURE a_2_in_2_out(param_one INT(11), param_two INT(11), OUT param_one INT(11), OUT param_two INT(11))
BEGIN
  SELECT param_one, param_two INTO return_one, return_two\;
END;

-- function a_func_0_in
CREATE FUNCTION a_func_0_in() RETURNS INT(11)
BEGIN
  RETURN 10\;
END;

-- function a_func_1_in
CREATE FUNCTION a_func_1_in(a_param INT(11)) RETURNS INT(11)
BEGIN
  RETURN a_param\;
END;

-- function a_func_2_in
CREATE FUNCTION a_func_2_in(param_one INT(11), param_two INT(11)) RETURNS INT(11)
BEGIN
  RETURN param_one + param_two\;
END;
