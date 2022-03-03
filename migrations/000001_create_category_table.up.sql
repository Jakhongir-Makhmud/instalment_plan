CREATE TABLE IF NOT EXISTS product_category (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(128) NOT NULL,
    category_percent DECIMAL(4,2) NOT NULL
);

INSERT INTO product_category (category_name,category_percent) values 
('phone',3),
('laptop',4),
('tv',5);