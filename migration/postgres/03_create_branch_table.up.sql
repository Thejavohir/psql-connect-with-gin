CREATE TABLE IF NOT EXISTS branch(
    id uuid primary key,
    name varchar(60) not null,
    address varchar,
    phone_number varchar(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);



CREATE TABLE IF NOT EXISTS branch_product_relation(
    product_id UUID REFERENCES product(id),
    branch_id UUID REFERENCES branch(id)
);

CREATE UNIQUE INDEX ON branch_product_relation(product_id, branch_id);
