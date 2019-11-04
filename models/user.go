package models

import (
	//"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/azurramas/food_ordering/services"
)


//User ->
type User struct {
	ID              int64                     `db:"id" json:"id"`
	Username        string                    `db:"username" json:"username"`
	Password        string                    `db:"password" json:"password"`
}

// Find -> Checks if user is lggedin
func (u *User) Find() (bool, error) {
	var user User
	db := services.SQLDbAccess.GetSQLDB()

	err := db.Get(&user, "SELECT * FROM users WHERE username = ?", u.Username)
	if err != nil {

		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if err == nil {
		u.ID = user.ID

		return true, err
	}
	return false, err

}

//Create ->
func (u *User) Create() error {
	//Get current DBAccess
	db := services.SQLDbAccess.GetSQLDB()

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)

	user, err := tx.Exec("INSERT INTO users(username, password) values(?,?)", u.Username, bytes)

	//If err -> Rollback
	if err != nil {
		tx.Rollback()
		return err
	}

	u.ID, err = user.LastInsertId()
	if err != nil {
		return err
	}

	//Commiting Transactions -> success
	tx.Commit()

	return nil
}
