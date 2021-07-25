-- SQL Schema template for the northwind schema.
-- Generated on Sun Jul 25 07:10:16 WIB 2021 by xo.

CREATE TABLE categories (
    category_id SMALLINT(6) NOT NULL,
    category_name VARCHAR(15) NOT NULL,
    description TEXT,
    picture BLOB,
    PRIMARY KEY (category_id)
) ENGINE=InnoDB;


CREATE TABLE customer_customer_demo (
    customer_id CHAR(255) NOT NULL REFERENCES customers (customer_id),
    customer_type_id CHAR(255) NOT NULL REFERENCES customer_demographics (customer_type_id),
    PRIMARY KEY (customer_id, customer_type_id)
) ENGINE=InnoDB;

CREATE INDEX customer_type_id ON customer_customer_demo (customer_type_id);

CREATE TABLE customer_demographics (
    customer_type_id CHAR(255) NOT NULL,
    customer_desc TEXT,
    PRIMARY KEY (customer_type_id)
) ENGINE=InnoDB;


CREATE TABLE customers (
    customer_id CHAR(255) NOT NULL,
    company_name VARCHAR(40) NOT NULL,
    contact_name VARCHAR(30),
    contact_title VARCHAR(30),
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    phone VARCHAR(24),
    fax VARCHAR(24),
    PRIMARY KEY (customer_id)
) ENGINE=InnoDB;


CREATE TABLE employee_territories (
    employee_id SMALLINT(6) NOT NULL REFERENCES employees (employee_id),
    territory_id VARCHAR(20) NOT NULL REFERENCES territories (territory_id),
    PRIMARY KEY (employee_id, territory_id)
) ENGINE=InnoDB;

CREATE INDEX territory_id ON employee_territories (territory_id);

CREATE TABLE employees (
    employee_id SMALLINT(6) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    first_name VARCHAR(10) NOT NULL,
    title VARCHAR(30),
    title_of_courtesy VARCHAR(25),
    birth_date DATE,
    hire_date DATE,
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    home_phone VARCHAR(24),
    extension VARCHAR(4),
    photo BLOB,
    notes TEXT,
    reports_to SMALLINT(6) REFERENCES employees (reports_to),
    photo_path VARCHAR(255),
    PRIMARY KEY (employee_id)
) ENGINE=InnoDB;

CREATE INDEX reports_to ON employees (reports_to);

CREATE TABLE order_details (
    order_id SMALLINT(6) NOT NULL REFERENCES orders (order_id),
    product_id SMALLINT(6) NOT NULL REFERENCES products (product_id),
    unit_price FLOAT NOT NULL,
    quantity SMALLINT(6) NOT NULL,
    discount FLOAT NOT NULL,
    PRIMARY KEY (order_id, product_id)
) ENGINE=InnoDB;

CREATE INDEX product_id ON order_details (product_id);

CREATE TABLE orders (
    order_id SMALLINT(6) NOT NULL,
    customer_id CHAR(255) REFERENCES customers (customer_id),
    employee_id SMALLINT(6) REFERENCES employees (employee_id),
    order_date DATE,
    required_date DATE,
    shipped_date DATE,
    ship_via SMALLINT(6) REFERENCES shippers (ship_via),
    freight FLOAT,
    ship_name VARCHAR(40),
    ship_address VARCHAR(60),
    ship_city VARCHAR(15),
    ship_region VARCHAR(15),
    ship_postal_code VARCHAR(10),
    ship_country VARCHAR(15),
    PRIMARY KEY (order_id)
) ENGINE=InnoDB;

CREATE INDEX customer_id ON orders (customer_id);
CREATE INDEX employee_id ON orders (employee_id);
CREATE INDEX ship_via ON orders (ship_via);

CREATE TABLE products (
    product_id SMALLINT(6) NOT NULL,
    product_name VARCHAR(40) NOT NULL,
    supplier_id SMALLINT(6) REFERENCES suppliers (supplier_id),
    category_id SMALLINT(6) REFERENCES categories (category_id),
    quantity_per_unit VARCHAR(20),
    unit_price FLOAT,
    units_in_stock SMALLINT(6),
    units_on_order SMALLINT(6),
    reorder_level SMALLINT(6),
    discontinued INT(11) NOT NULL,
    PRIMARY KEY (product_id)
) ENGINE=InnoDB;

CREATE INDEX category_id ON products (category_id);
CREATE INDEX supplier_id ON products (supplier_id);

CREATE TABLE region (
    region_id SMALLINT(6) NOT NULL,
    region_description CHAR(255) NOT NULL,
    PRIMARY KEY (region_id)
) ENGINE=InnoDB;


CREATE TABLE shippers (
    shipper_id SMALLINT(6) NOT NULL,
    company_name VARCHAR(40) NOT NULL,
    phone VARCHAR(24),
    PRIMARY KEY (shipper_id)
) ENGINE=InnoDB;


CREATE TABLE suppliers (
    supplier_id SMALLINT(6) NOT NULL,
    company_name VARCHAR(40) NOT NULL,
    contact_name VARCHAR(30),
    contact_title VARCHAR(30),
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    phone VARCHAR(24),
    fax VARCHAR(24),
    homepage TEXT,
    PRIMARY KEY (supplier_id)
) ENGINE=InnoDB;


CREATE TABLE territories (
    territory_id VARCHAR(20) NOT NULL,
    territory_description CHAR(255) NOT NULL,
    region_id SMALLINT(6) NOT NULL REFERENCES region (region_id),
    PRIMARY KEY (territory_id)
) ENGINE=InnoDB;

CREATE INDEX region_id ON territories (region_id);

CREATE TABLE us_states (
    state_id SMALLINT(6) NOT NULL,
    state_name VARCHAR(100),
    state_abbr VARCHAR(2),
    state_region VARCHAR(50),
    PRIMARY KEY (state_id)
) ENGINE=InnoDB;


