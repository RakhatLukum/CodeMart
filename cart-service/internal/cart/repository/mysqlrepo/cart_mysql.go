package mysqlrepo

import "database/sql"

// No import changes needed here for now.

type repo struct{ db *sql.DB }

func New(db *sql.DB) *repo { return &repo{db: db} }

func (r *repo) Add(userID, productID int64) error {
	_, err := r.db.Exec(`INSERT INTO carts(user_id,product_id) VALUES(?,?)`, userID, productID)
	return err
}
func (r *repo) Remove(userID, productID int64) error {
	_, err := r.db.Exec(`DELETE FROM carts WHERE user_id=? AND product_id=? LIMIT 1`, userID, productID)
	return err
}
func (r *repo) List(userID int64) ([]int64, error) {
	rows, err := r.db.Query(`SELECT product_id FROM carts WHERE user_id=?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (r *repo) Clear(userID int64) error {
	_, err := r.db.Exec(`DELETE FROM carts WHERE user_id=?`, userID)
	return err
}
func (r *repo) Count(userID int64) (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM carts WHERE user_id=?`, userID).Scan(&count)
	return count, err
}
func (r *repo) Has(userID, productID int64) (bool, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM carts WHERE user_id=? AND product_id=?`, userID, productID).Scan(&count)
	return count > 0, err
}
func (r *repo) Replace(userID int64, productIDs []int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM carts WHERE user_id=?`, userID); err != nil {
		tx.Rollback()
		return err
	}
	for _, pid := range productIDs {
		if _, err := tx.Exec(`INSERT INTO carts(user_id,product_id) VALUES(?,?)`, userID, pid); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
func (r *repo) AddMultiple(userID int64, productIDs []int64) error {
	for _, pid := range productIDs {
		if _, err := r.db.Exec(`INSERT INTO carts(user_id,product_id) VALUES(?,?)`, userID, pid); err != nil {
			return err
		}
	}
	return nil
}
func (r *repo) RemoveMultiple(userID int64, productIDs []int64) error {
	for _, pid := range productIDs {
		if _, err := r.db.Exec(`DELETE FROM carts WHERE user_id=? AND product_id=? LIMIT 1`, userID, pid); err != nil {
			return err
		}
	}
	return nil
}
