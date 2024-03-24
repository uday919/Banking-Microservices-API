package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uday919/Banking-Microservices-API/domain"
	"github.com/uday919/Banking-Microservices-API/service"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/greet", greet)

	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
