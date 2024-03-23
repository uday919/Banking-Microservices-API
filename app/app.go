package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)

	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
