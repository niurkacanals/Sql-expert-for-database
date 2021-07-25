-- SQL Schema template for the northwind schema.
-- Generated on Sun Jul 25 07:10:29 WIB 2021 by xo.

CREATE TABLE categories (
    category_id SMALLINT(5) NOT NULL,
    category_name VARCHAR(15) NOT NULL,
    description TEXT,
    picture VARBINARY,
    CONSTRAINT categories_pkey PRIMARY KEY (category_id)
);


CREATE TABLE customer_customer_demo (
    customer_id CHAR(255) NOT NULL CONSTRAINT customer_customer_demo_customer_id_fkey REFERENCES customers (customer_id),
    customer_type_id CHAR(255) NOT NULL CONSTRAINT customer_customer_demo_customer_type_id_fkey REFERENCES customer_demographics (customer_type_id),
    CONSTRAINT customer_customer_demo_pkey PRIMARY KEY (customer_id, customer_type_id)
);


CREATE TABLE customer_demographics (
    customer_type_id CHAR(255) NOT NULL,
    customer_desc TEXT,
    CONSTRAINT customer_demographics_pkey PRIMARY KEY (customer_type_id)
);


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
    CONSTRAINT customers_pkey PRIMARY KEY (customer_id)
);


CREATE TABLE employee_territories (
    employee_id SMALLINT(5) NOT NULL CONSTRAINT employee_territories_employee_id_fkey REFERENCES employees (employee_id),
    territory_id VARCHAR(20) NOT NULL CONSTRAINT employee_territories_territory_id_fkey REFERENCES territories (territory_id),
    CONSTRAINT employee_territories_pkey PRIMARY KEY (employee_id, territory_id)
);


CREATE TABLE employees (
    employee_id SMALLINT(5) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    first_name VARCHAR(10) NOT NULL,
    title VARCHAR(30),
    title_of_courtesy VARCHAR(25),
    birth_date DATE(10),
    hire_date DATE(10),
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    home_phone VARCHAR(24),
    extension VARCHAR(4),
    photo VARBINARY,
    notes TEXT,
    reports_to SMALLINT(5) CONSTRAINT employees_reports_to_fkey REFERENCES employees (reports_to),
    photo_path VARCHAR(255),
    CONSTRAINT employees_pkey PRIMARY KEY (employee_id)
);


CREATE TABLE order_details (
    order_id SMALLINT(5) NOT NULL CONSTRAINT order_details_order_id_fkey REFERENCES orders (order_id),
    product_id SMALLINT(5) NOT NULL CONSTRAINT order_details_product_id_fkey REFERENCES products (product_id),
    unit_price REAL(24) NOT NULL,
    quantity SMALLINT(5) NOT NULL,
    discount REAL(24) NOT NULL,
    CONSTRAINT order_details_pkey PRIMARY KEY (order_id, product_id)
);


CREATE TABLE orders (
    order_id SMALLINT(5) NOT NULL,
    customer_id CHAR(255) CONSTRAINT orders_customer_id_fkey REFERENCES customers (customer_id),
    employee_id SMALLINT(5) CONSTRAINT orders_employee_id_fkey REFERENCES employees (employee_id),
    order_date DATE(10),
    required_date DATE(10),
    shipped_date DATE(10),
    freight REAL(24),
    ship_name VARCHAR(40),
    ship_address VARCHAR(60),
    ship_city VARCHAR(15),
    ship_region VARCHAR(15),
    ship_postal_code VARCHAR(10),
    ship_country VARCHAR(15),
    CONSTRAINT orders_pkey PRIMARY KEY (order_id)
);


CREATE TABLE products (
    product_id SMALLINT(5) NOT NULL,
    product_name VARCHAR(40) NOT NULL,
    supplier_id SMALLINT(5) CONSTRAINT products_suplier_id_fkey REFERENCES suppliers (supplier_id),
    category_id SMALLINT(5) CONSTRAINT products_category_id_fkey REFERENCES categories (category_id),
    quantity_per_unit VARCHAR(20),
    unit_price REAL(24),
    units_in_stock SMALLINT(5),
    units_on_order SMALLINT(5),
    reorder_level SMALLINT(5),
    discontinued INT(10) NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (product_id)
);


CREATE TABLE region (
    region_id SMALLINT(5) NOT NULL,
    region_description CHAR(255) NOT NULL,
    CONSTRAINT regions_pkey PRIMARY KEY (region_id)
);


CREATE TABLE shippers (
    shipper_id SMALLINT(5) NOT NULL,
    company_name VARCHAR(40) NOT NULL,
    phone VARCHAR(24),
    CONSTRAINT shippers_pkey PRIMARY KEY (shipper_id)
);


CREATE TABLE suppliers (
    supplier_id SMALLINT(5) NOT NULL,
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
    CONSTRAINT suppliers_pkey PRIMARY KEY (supplier_id)
);


CREATE TABLE territories (
    territory_id VARCHAR(20) NOT NULL,
    territory_description CHAR(255) NOT NULL,
    region_id SMALLINT(5) NOT NULL CONSTRAINT territories_region_id_fkey REFERENCES region (region_id),
    CONSTRAINT territories_pkey PRIMARY KEY (territory_id)
);


CREATE TABLE us_states (
    state_id SMALLINT(5) NOT NULL,
    state_name VARCHAR(100),
    state_abbr VARCHAR(2),
    state_region VARCHAR(50),
    CONSTRAINT us_states_pkey PRIMARY KEY (state_id)
);


