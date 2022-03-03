CREATE TABLE IF NOT EXISTS product_installment (
    id UUID PRIMARY KEY,
    product_name VARCHAR(128) NOT NULL,
    category_id INT NOT NULL,
    product_price DECIMAL(8,4) NOT NULL,
    addition_price DECIMAL(8,4) NOT NULL,
    installment_range INT NOT NULL,
    paid_part DECIMAL(8,4) DEFAULT 0
);