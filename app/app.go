package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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
	dbClient := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	// wiring
	// ch := CustomersHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomersHandlers{service: service.NewCustomerService(customerRepositoryDb)}

	router.HandleFunc("/customers", ch.getAllCoustumers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPass,
		dbAddr,
		dbPort,
		dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
