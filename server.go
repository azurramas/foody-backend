package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/azurramas/food_ordering/controllers"
	"github.com/azurramas/food_ordering/middleware"
	"github.com/azurramas/food_ordering/services"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	requests controllers.RequestController
	users    controllers.UserControler
	provider middleware.Provider
	mdlw     middleware.Middlewares
	orders   controllers.OrdersController
)

var err error

func main() {
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

	//Requests CRUD
	//OID - OrderID
	r1.HandleFunc("/request/{id:[0-9]+}", requests.CreateByOID).Methods("POST")
	r1.HandleFunc("/requests/{id}", requests.ListByParam).Methods("GET")

	//Orders CRUD
	r1.HandleFunc("/order", orders.Create).Methods("POST")
	r1.HandleFunc("/orders", orders.ListAll).Methods("GET")
	r1.HandleFunc("/user/order", orders.ListByUser).Methods("GET")
	r1.HandleFunc("/order/{id:[0-9]+}", orders.DeleteOrder).Methods("DELETE")

	//WebSocket Endpoint
	r1.HandleFunc("/ws/{id:[0-9]+}", controllers.WsHandler)

	n.UseHandler(r1)

	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8010", n))

}
