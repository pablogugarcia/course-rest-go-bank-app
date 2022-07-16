package app

import (
	"encoding/json"
	"net/http"

	dto "github.com/pablogugarcia/banking/dtos"
	"github.com/pablogugarcia/banking/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (th TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		transaction, err := th.service.NewTransaction(&request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusCreated, transaction)
		}
	}
}
