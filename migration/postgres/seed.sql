BEGIN;

-- Seed categories
INSERT INTO category (name) VALUES
('Electronics'),
('Books'),
('Clothing');

-- Seed products
INSERT INTO product (id, name, description, image_url, category_code)
SELECT gen_random_uuid(), 'Smartphone', 'A modern smartphone', 'https://example.com/pic1.png', code FROM category WHERE name = 'Electronics'
UNION ALL
SELECT gen_random_uuid(), 'Novel', 'A bestselling novel', 'https://example.com/pic2.png', code FROM category WHERE name = 'Books'
UNION ALL
SELECT gen_random_uuid(), 'T-Shirt', 'Cotton t-shirt', 'https://example.com/pic3.png', code FROM category WHERE name = 'Clothing';

COMMIT;