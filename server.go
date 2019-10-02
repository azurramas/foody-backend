package main

import (
	"log"
	"net/http"
	"github.com/azurramas/food_ordering/services"
	"github.com/azurramas/food_ordering/middleware"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/azurramas/food_ordering/controllers"
)

var (
	requests	 controllers.RequestController
	users        controllers.UserControler
	provider     middleware.Provider
	mdlw         middleware.Middlewares
)

var err error

func main(){
	//Making initial DB Access
	services.GetDBAccess("dev", "conf/conf.yaml")

	//Router
	r1 := mux.NewRouter()

	// NEGRONI MIDDLEWARE
	n := negroni.Classic()

	n.Use(negroni.HandlerFunc(mdlw.CORS))
	n.Use(negroni.HandlerFunc(mdlw.Preflight))
	n.Use(negroni.HandlerFunc(provider.BasicAuth))
	
	//Users CRUD
	r1.HandleFunc("/user", users.Create).Methods("POST")
	r1.HandleFunc("/login", users.Login).Methods("POST")
	// r1.HandleFunc("/user/{id:[0-9]+}", users.Get).Methods("GET")
	// r1.HandleFunc("/users", users.List).Methods("GET")

	//Requests CRUD 
	//OID - OrderID
	r1.HandleFunc("/request/{id:[0-9]+}", requests.CreateByOID).Methods("POST")
	r1.HandleFunc("/requests/{id}", requests.ListByParam).Methods("GET")
	
	n.UseHandler(r1)

	log.Fatal(http.ListenAndServe(":8010", n))

}