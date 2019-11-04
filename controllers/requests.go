package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/azurramas/food_ordering/models"
	"github.com/azurramas/food_ordering/services"
)

// RequestController ->
type RequestController struct{}


//CreateByOID ->
func (rc RequestController) CreateByOID(w http.ResponseWriter, r *http.Request) {
	var request models.Request

	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, 400)
		return
	}

	err = request.Create(id)
	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, 400)
		return
	}

	services.WriteJSON(w, request, 200)

}

//ListByParam ->
func (rc RequestController) ListByParam(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	id := params["id"]
	
	requests, err := models.ListRequests(id)
	if err != nil {
		fmt.Println(err)
		services.WriteJSON(w, err, 400)
		return
	}
	
	if len(requests) == 0 {
		services.WriteJSON(w, "No requests found", http.StatusBadRequest)
		return
	}

	services.WriteJSON(w, requests, 200)
}
