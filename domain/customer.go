package domain

import "github.com/pablogugarcia/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppErr)
	ById(id string) (*Customer, *errs.AppErr)
}
