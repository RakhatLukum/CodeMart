CREATE TABLE IF NOT EXISTS user_view_counts (
    user_id INT PRIMARY KEY,
    total_views INT NOT NULL DEFAULT 0,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
