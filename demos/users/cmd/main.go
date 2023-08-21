package main

import (
	"log"
	"net/http"

	"github.com/jwkidd3/golang-microservices/demos/users/internal/routes"
)

func main() {
	r := routes.Handlers()

	err := http.ListenAndServe(":5050", r)
	if err != nil {
		log.Fatal(err)
	}
}
