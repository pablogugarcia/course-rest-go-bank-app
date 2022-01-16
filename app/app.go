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
	ch := CustomersHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCoustumers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
