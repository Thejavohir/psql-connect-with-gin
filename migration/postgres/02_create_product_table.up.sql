CREATE TABLE IF NOT EXISTS product(
    id uuid primary key,
    name varchar(60) not null,
    price numeric not null,
    category_id uuid REFERENCES category(id),
    barcode varchar(60) not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
