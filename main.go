package main

import (
	"log"
	"net/http"

	"github.com/zemirco/redirect/router"
)

func main() {
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
