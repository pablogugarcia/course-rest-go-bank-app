package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pablogugarcia/banking/errs"
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

func (ch *CustomersHandlers) getAllCoustumers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	if status != "" {
		if !isValidStatus(status) {
			err := errs.NewNotFoundError("bad request")
			writeResponse(w, err.Code, err.AsMessage())
			return
		} else {
			var err error
			status, err = parseStatus(status)

			if err != nil {
				err := errs.NewUnexpectedError("internal server error")
				writeResponse(w, err.Code, err.AsMessage())
			}
		}
	}

	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomersHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)["customer_id"]

	customer, err := ch.service.GetCustomer(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func isValidStatus(status string) bool {
	return status == "active" || status == "inactive"
}

func parseStatus(status string) (string, error) {
	if status == "active" {
		return "1", nil
	} else if status == "inactive" {
		return "0", nil
	} else {
		return "", errors.New("bad status param")
	}
}
