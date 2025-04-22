-- Users
INSERT INTO users (email, password) VALUES
    ('alice@example.com', '123456'),
    ('bob@example.com', 'abcdef');

-- Products
INSERT INTO products (name, price, tags) VALUES
    ('Milk', 500.00, '["halal"]'),
    ('Tofu', 700.00, '["vegan"]'),
    ('Honey', 900.00, '[]');

-- Carts
INSERT INTO carts (user_id, product_id) VALUES
    (1, 1),
    (1, 3),
    (2, 2);

-- Views
INSERT INTO views (user_id, product_id, timestamp) VALUES
    (1, 1, '2025-04-20 17:38:45'),
    (1, 1, '2025-04-20 17:38:45'),
    (2, 2, '2025-04-20 17:38:45'),
    (2, 1, '2025-04-20 17:38:45'),
    (1, 3, '2025-04-20 17:38:45');
