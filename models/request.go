package models

import (
	//"fmt"
	"github.com/azurramas/food_ordering/services"
)

//Request ->
type Request struct {
	UserName			string		`db:"user_name" json:"user_name"`
	RequestContent     	string      `db:"request_content" json:"request_content"`
	OrderID				string      `db:"order_id" json:"order_id"`
}


//Create ->
func (r *Request) Create(id string) error {
	
	db := services.SQLDbAccess.GetSQLDB()
	
	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO requests( user_name, request_content, order_id ) values(?,?,?)", r.UserName, r.RequestContent, id )

	if err != nil {
		tx.Rollback()
		return err
	}
	r.OrderID = id
	
	tx.Commit()
	return nil
}

//ListRequests ->
func ListRequests(id string) ([]Request, error){
	var requests []Request

	db := services.SQLDbAccess.GetSQLDB()

	err := db.Select(&requests, "SELECT * FROM requests WHERE order_id=? OR user_name=?", id, id)
	
	return requests, err
	
}


