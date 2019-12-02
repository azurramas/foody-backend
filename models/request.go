package models

import (
	"strconv"
	"errors"
	"database/sql"
	//"fmt"
	"github.com/azurramas/food_ordering/services"
)

//Request ->
type Request struct {
	ID					int64      	`db:"id" json:"id"`	
	UserName			string		`db:"user_name" json:"user_name"`
	RequestContent     	string      `db:"request_content" json:"request_content"`
	OrderID				int64     `db:"order_id" json:"order_id"`
}


//Create ->
func (r *Request) Create(strid string) error {

	id, err := strconv.ParseInt(strid, 10, 64)

	OID, err := GetOrderByID(id)

	if (Order{}) == OID  {
		err = errors.New("Order does not exist")
		return err
	}
	
	db := services.SQLDbAccess.GetSQLDB()
	
	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	request, err := tx.Exec("INSERT INTO requests( user_name, request_content, order_id ) values(?,?,?)", r.UserName, r.RequestContent, id )

	if err != nil {
		tx.Rollback()
		return err
	}

	r.ID, err = request.LastInsertId()
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
	
//DeleteRequestsForOrder ->
func DeleteRequestsForOrder(id int64, tx *sql.Tx) (*sql.Tx, error) {
	_, err := tx.Exec("DELETE FROM requests WHERE order_id=?", id)

	return tx, err
}


