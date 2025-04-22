package mysqlrepo

import "database/sql"

type repo struct{ db *sql.DB }

func New(db *sql.DB)*repo{ return &repo{db:db}}

func (r *repo) Add(userID,productID int64) error{
    _,err:=r.db.Exec(`INSERT INTO carts(user_id,product_id) VALUES(?,?)`,userID,productID)
    return err
}
func (r *repo) Remove(userID,productID int64) error{
    _,err:=r.db.Exec(`DELETE FROM carts WHERE user_id=? AND product_id=? LIMIT 1`,userID,productID)
    return err
}
func (r *repo) List(userID int64)([]int64,error){
    rows,err:=r.db.Query(`SELECT product_id FROM carts WHERE user_id=?`,userID)
    if err!=nil{return nil,err}
    defer rows.Close()
    var ids []int64
    for rows.Next(){
        var id int64
        if err:=rows.Scan(&id);err!=nil{return nil,err}
        ids=append(ids,id)
    }
    return ids,nil
}
