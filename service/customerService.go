package service

import (
	"github.com/pablogugarcia/banking/domain"
	dto "github.com/pablogugarcia/banking/dtos"
	"github.com/pablogugarcia/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppErr)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppErr) {
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	dtoList := make([]dto.CustomerResponse, 0)
	for _, customer := range c {
		dtoList = append(dtoList, customer.ToDto())
	}
	return dtoList, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppErr) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
