CREATE TABLE IF NOT EXISTS product_category (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(128) NOT NULL,
    category_percent DECIMAL(4,2) NOT NULL
);