package service

import (
	"github.com/pablogugarcia/banking/domain"
	"github.com/pablogugarcia/banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, *errs.AppErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppErr) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
