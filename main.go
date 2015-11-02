package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(index)
	r.Methods("POST").Path("/login").HandlerFunc(login)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// GET /
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there!")
}

// POST /login
func login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
