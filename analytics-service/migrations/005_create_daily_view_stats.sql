CREATE TABLE IF NOT EXISTS daily_view_stats (
    product_id INT NOT NULL,
    view_date DATE NOT NULL,
    view_count INT NOT NULL DEFAULT 0,
    PRIMARY KEY (product_id, view_date),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
