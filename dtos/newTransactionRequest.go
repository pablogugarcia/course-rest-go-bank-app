package dto

import (
	"strings"

	"github.com/pablogugarcia/banking/errs"
)

type NewTransactionRequest struct {
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	AccountId   string  `json:"account_id"`
	OpeningDate string  `json:"opening_date"`
}

func (r NewTransactionRequest) Validate() *errs.AppErr {
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be negative")
	}
	if strings.ToLower(r.Type) != "withdrawl" && strings.ToLower(r.Type) != "deposit" {
		return errs.NewValidationError("Incorrect transaction type")
	}
	return nil
}
