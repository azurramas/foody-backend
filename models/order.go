package models

import (
	"strconv"
	//"fmt"
	//"errors"
	"github.com/azurramas/food_ordering/services"
)

//Order ->
type Order struct {
	UserName		string		`db:"user_name" json:"user_name"`
	UserID			int64		`db:"user_id" json:"user_id"`
	ID				int64      	`db:"id" json:"id"`
	Restaurant     	string      `db:"restaurant" json:"restaurant"`
	Comment 		string		`db:"comment" json:"comment"`
}

//Create ->
func (o *Order) Create(id int64, username string) error {
	//Get current DBAccess
	db := services.SQLDbAccess.GetSQLDB()

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	order, err := tx.Exec("INSERT INTO orders(user_name, user_id, restaurant, comment) values(?,?,?,?)",username, id, o.Restaurant, o.Comment)

	//If err -> Rollback
	if err != nil {
		tx.Rollback()
		return err
	}
	o.UserID = id
	o.ID, err = order.LastInsertId()
	if err != nil {
		return err
	}

	//Commiting Transactions -> success
	tx.Commit()

	return nil
}

//ListAllOrders ->
func ListAllOrders() ([]Order, error){
	var orders []Order

	//Get current DBAccess
	db := services.SQLDbAccess.GetSQLDB()

	err := db.Select(&orders, "SELECT * FROM orders")
	return orders, err
}

//ListOrdersByUID ->
func ListOrdersByUID(uid int64) ([]Order, error){
	var orders []Order

	//Get current DBAccess
	db := services.SQLDbAccess.GetSQLDB()

	err := db.Select(&orders, "SELECT * FROM orders WHERE user_id=?", uid)
	return orders, err
}

//Delete ->
func (o *Order) Delete(strid string, uID int64) error {

	id, err := strconv.ParseInt(strid, 10, 64)
	//Get current DBAccess
	db := services.SQLDbAccess.GetSQLDB()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM orders WHERE id=? AND user_id=?", id, uID)
	if err != nil {
		tx.Rollback()
		return err
	}
	
	DeleteRequestsForOrder(id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//GetOrderByID ->
func GetOrderByID(id int64) (Order, error){
	var order Order

	//Get current DBAccess
	db := services.SQLDbAccess.GetSQLDB()

	err := db.Get(&order, "SELECT * FROM orders WHERE id=?", id)
	return order, err
}