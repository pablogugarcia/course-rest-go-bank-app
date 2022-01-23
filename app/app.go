package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pablogugarcia/banking/domain"
	"github.com/pablogugarcia/banking/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	// ch := CustomersHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomersHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCoustumers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
