package middleware

import (
	// "fmt"
	"net/http"
)

// CORS ->
func (m Middlewares) CORS(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	//fmt.Println("CORS")
	// CORS support for Preflighted requests
	// TODO: replace * for client origin domain

	//res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	// res.Header().Set("Access-Control-Max-Age", "1728000")
	//fmt.Println("CORS2")
	next(res, req)

}

//Preflight -> 
func (m Middlewares) Preflight(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// NOTE: publicly accessable routes
	origin := req.Header.Get("Origin")
	res.Header().Set("Access-Control-Allow-Origin", origin)
	if req.Method == "OPTIONS" {
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")

		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return
	}
	 
	next(res, req)
}
