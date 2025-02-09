CREATE TABLE IF NOT EXISTS market(
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    address VARCHAR,
    phone_number VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);


CREATE TABLE IF NOT EXISTS market_product_relation(
    product_id UUID REFERENCES product(id),
    market_id UUID REFERENCES market(id),
    UNIQUE (product_id, market_id)
);