package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/azurramas/food_ordering/models"
	"github.com/azurramas/food_ordering/services"
)

// UserControler ->
type UserControler struct{}

//TryAuthenticate ->function for basic authentification
func (uc UserControler) TryAuthenticate(w http.ResponseWriter, r *http.Request, u *models.User) (bool, error) {

	var hasaccess bool
	hasaccess, err := u.Find()

	return hasaccess, err

}

//Login ->
func (uc UserControler) Login(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value("user").(models.User)

	services.WriteJSON(w, u, 200)
}


//Create ->
func (uc UserControler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, 400)
		return
	}

	if user.Username == "" || len(user.Username) < 3 || user.Password == "" || len(user.Password) < 3 {
		errMessage := "Username/Password can not be empty or less then 3 charactes!!"
		fmt.Println(errMessage)
		services.WriteJSON(w, errMessage, 400)
		return
	}

	err = user.Create()
	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, 400)
		return
	}

	services.WriteJSON(w, user, 200)

}
