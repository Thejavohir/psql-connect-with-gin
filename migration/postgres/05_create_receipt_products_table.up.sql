
CREATE TABLE receipt_products(
    id uuid primary key,
    name varchar(60) not null,
    receipt_id uuid REFERENCES receipt(id),
    category_id uuid REFERENCES category(id),
    barcode varchar(60) not null,
    quantity integer not null DEFAULT 0,
    price numeric not null,
    total_price numeric DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
