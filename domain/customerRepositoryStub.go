package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"10001", "Ashley", "New York", "4000", "10/10/98", "1"},
		{"10002", "Rob", "Lima", "200", "15/8/98", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
