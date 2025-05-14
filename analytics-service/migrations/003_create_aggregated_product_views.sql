CREATE TABLE IF NOT EXISTS aggregated_product_views (
    product_id INT PRIMARY KEY,
    total_views INT NOT NULL DEFAULT 0,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id)
);
