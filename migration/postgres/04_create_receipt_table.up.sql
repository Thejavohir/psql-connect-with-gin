
CREATE TABLE IF NOT EXISTS receipt(
    id uuid primary key,
    receipt_id varchar,
    branch_id uuid REFERENCES branch(id),
    date_time TIMESTAMP,
    status varchar DEFAULT 'in_process',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP 
);
