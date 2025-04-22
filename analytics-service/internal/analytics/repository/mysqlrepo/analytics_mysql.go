package mysqlrepo

import "database/sql"

type repo struct{ db *sql.DB }

func New(db *sql.DB)*repo{ return &repo{db:db}}

func (r *repo) Inc(productID int64) error{
    _,err:=r.db.Exec(`INSERT INTO product_views(product_id,view_count) VALUES(?,1) ON DUPLICATE KEY UPDATE view_count=view_count+1`,productID)
    return err
}
func (r *repo) Top(limit int)([]int64,error){
    rows,err:=r.db.Query(`SELECT product_id FROM product_views ORDER BY view_count DESC LIMIT ?`,limit)
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
