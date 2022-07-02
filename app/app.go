package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pablogugarcia/banking/domain"
	"github.com/pablogugarcia/banking/service"
)

func sanityCheck() {
	envs := []string{
		"DB_USER",
		"DB_PASSWORD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
		"SERVER_PORT",
		"SERVER_ADDRESS"}
	for _, env := range envs {
		if os.Getenv(env) == "" {
			log.Fatalf("Missing enviroment variable %s", env)
		}
	}

}

func Start() {
	godotenv.Load()
	sanityCheck()
	router := mux.NewRouter()

	// wiring
	// ch := CustomersHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomersHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCoustumers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
