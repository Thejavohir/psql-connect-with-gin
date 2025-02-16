
CREATE TABLE remainder(
    id uuid primary key,
    name varchar(60) not null,
    category_id uuid REFERENCES category(id),
    branch_id uuid REFERENCES branch(id), 
    barcode varchar(60) not null,
    quantity integer not null DEFAULT 0,
    price numeric not null,
    total_price numeric DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);