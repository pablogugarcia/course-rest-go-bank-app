package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/pablogugarcia/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomersHandlers struct {
	service service.CustomerService
}

func (ch *CustomersHandlers) getAllCoustumers(rw http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Alish", City: "Cba", Zipcode: "5000"},
	// 	{Name: "Castor", City: "Cba", Zipcode: "5000"},
	// }
	customers, _ := ch.service.GetAllCustomers()
	if r.Header.Get("Content-type") == "application/xml" {
		rw.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)

	} else {
		rw.Header().Add("Content-type", "application/json")
		json.NewEncoder(rw).Encode(customers)

	}

}
