package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/azurramas/food_ordering/models"
	"github.com/azurramas/food_ordering/services"
)

//OrdersController ->
type OrdersController struct{}

//Create ->
func (oc OrdersController) Create(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value("user").(models.User)

	var order models.Order

	order.UserName = u.Username

	err := json.NewDecoder(r.Body).Decode(&order)
	
	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, http.StatusBadRequest)
		return
	}

	err = order.Create(u.ID, u.Username)
	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, http.StatusBadRequest)
		return
	}
	
	services.WriteJSON(w, order, http.StatusOK)
}

//ListAll ->
func (oc OrdersController) ListAll(w http.ResponseWriter, r *http.Request) {

	orders, err := models.ListAllOrders()

	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, http.StatusBadRequest)
		return
	}

	services.WriteJSON(w, orders, http.StatusOK)
}


//ListByUser ->
func (oc OrdersController) ListByUser(w http.ResponseWriter, r *http.Request){
	u := r.Context().Value("user").(models.User)

	order, err := models.ListOrdersByUID(u.ID)

	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, http.StatusBadRequest)
		return
	}

	services.WriteJSON(w, order, http.StatusOK)
	
}

//DeleteOrder ->
func (oc OrdersController) DeleteOrder(w http.ResponseWriter, r *http.Request){
	var order models.Order

	params := mux.Vars(r)
	id := params["id"]

	u := r.Context().Value("user").(models.User)

	err := order.Delete(id, u.ID)

	if err != nil {
		fmt.Println(err)
		success := map[string]interface{}{"success": "false"}
		services.WriteJSON(w, success, http.StatusBadRequest)
		return
	}

	success := map[string]interface{}{"success": "true"}
	services.WriteJSON(w, success, http.StatusOK)

}



