package middleware

import (
	"context"
	//"errors"
	//"fmt"
	"net/http"
	"strings"
	"github.com/casbin/casbin"
	"github.com/azurramas/food_ordering/services"
	"github.com/azurramas/food_ordering/models"
	"github.com/azurramas/food_ordering/controllers"
)

// Middlewares ->
type Middlewares struct{}

// Provider ->
type Provider struct {
	rules *casbin.Enforcer
}

// BasicAuth ->
func (m *Provider) BasicAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if strings.Contains(r.RequestURI, "/request/") && r.Method == "POST"{
		next(w, r)
		return
	}

	if strings.Contains(r.RequestURI, "/orders") && r.Method == "GET" {
		next(w, r)
		return
	}

	if (r.RequestURI == "/user" && r.Method == "POST") {
		next(w, r)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		services.WriteJSON(w, "Authorization header format must be Basic ", http.StatusUnauthorized)
		return
	}

	user := models.User{Username: username, Password: password}

	userControler := controllers.UserControler{}

	var hasaccess bool
	hasaccess, r1 := userControler.TryAuthenticate(w, r, &user)

	if r1 == nil && hasaccess {

		c := r.Context()
		next(w, r.WithContext(context.WithValue(c, "user", user)))

	} else {
		services.WriteJSON(w, "Authentication failed", http.StatusUnauthorized)
	}
}