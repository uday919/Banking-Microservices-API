package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uday919/Banking-Microservices-API/service"
)

type Customer struct {
	Name    string `json:"full_name "xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}
type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Ashish", City: "New Delhi", Zipcode: "110093"},
	// 	{Name: "John Doe", City: "Los Angeles", Zipcode: "90210"},
	// }
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
