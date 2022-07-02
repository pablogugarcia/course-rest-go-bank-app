package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pablogugarcia/banking/errs"
	"github.com/pablogugarcia/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppErr) {
	var err error
	customers := make([]Customer, 0)

	if status == "0" || status == "1" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		d.client.Select(&customers, findAllSql, status)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		d.client.Select(&customers, findAllSql)
	}

	if err != nil {
		logger.Error("Error while quering customers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppErr) {
	findById := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer

	err := d.client.Get(&c, findById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while scaning customers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
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
	return CustomerRepositoryDb{client}
}
