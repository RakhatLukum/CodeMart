CREATE TABLE IF NOT EXISTS hourly_view_stats (
    product_id INT NOT NULL,
    view_hour DATETIME NOT NULL,
    view_count INT NOT NULL DEFAULT 0,
    PRIMARY KEY (product_id, view_hour),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
