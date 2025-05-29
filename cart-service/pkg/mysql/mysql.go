package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DSN      string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type DB struct {
	Conn *sql.DB
}

func NewDB(cfg Config) (*DB, error) {
	var dsn string
	if cfg.DSN != "" {
		dsn = cfg.DSN
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("mysql connection failed: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("mysql ping failed: %w", err)
	}

	return &DB{Conn: db}, nil
}

func (d *DB) Close() error {
	return d.Conn.Close()
}
