package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates new router with all http handlers
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(index)
	r.Methods("POST").Path("/login").HandlerFunc(login)
	return r
}

// GET /
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there!")
}

// POST /login
func login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
