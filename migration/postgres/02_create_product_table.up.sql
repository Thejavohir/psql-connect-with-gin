CREATE TABLE "product" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "category_id" UUID REFERENCES "category"("id") ON DELETE SET NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

