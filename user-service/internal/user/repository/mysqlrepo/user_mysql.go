package mysqlrepo

import (
	"CodeMart/user-service/internal/user/entity"
	"database/sql"
)

type repo struct{ db *sql.DB }

func New(db *sql.DB) *repo { return &repo{db: db} }

func (r *repo) Create(email, password string) (int64, error) {
	res, err := r.db.Exec(`INSERT INTO users(email,password) VALUES(?,?)`, email, password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func (r *repo) GetByEmailAndPassword(email, password string) (*entity.User, error) {
	var u entity.User
	err := r.db.QueryRow(`SELECT id,email,password FROM users WHERE email=? AND password=?`, email, password).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *repo) GetByID(id int64) (*entity.User, error) {
	var u entity.User
	err := r.db.QueryRow(`SELECT id,email,password FROM users WHERE id=?`, id).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
