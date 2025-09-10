BEGIN;

CREATE TABLE IF NOT EXISTS category(
    code SERIAL PRIMARY KEY,
    name VARCHAR(80) UNIQUE
);

CREATE TABLE IF NOT EXISTS product(
    id UUID PRIMARY KEY,
    name VARCHAR(80) NOT NULL,
    description TEXT NOT NULL,
    image_url TEXT DEFAULT '' NOT NULL,
    category_code INTEGER NOT NULL REFERENCES category(code),

    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at  TIMESTAMPTZ,
    deleted_at  TIMESTAMPTZ
);

CREATE INDEX idx_product_name ON product (name);

COMMIT;