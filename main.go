package main

import (
	"fmt"
	"log"
	"net/http"

	tx "github.com/agustin-sarasua/rs-transaction-api/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/transaction", tx.CreateTransactionEndpoint).Methods("POST")

	fmt.Println("Hello Transaction API")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
