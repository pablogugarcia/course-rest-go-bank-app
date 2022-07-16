package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/pablogugarcia/banking/errs"
	"github.com/pablogugarcia/banking/logger"
)

type TransactionRepository struct {
	client *sqlx.DB
}

func (t TransactionRepository) Save(tx *Transaction) (*Transaction, *errs.AppErr) {
	sqlInsert := "INSERT INTO transactions ( amount, type) values (?,?)"
	result, err := t.client.Exec(sqlInsert, &tx.Amount, &tx.TransactionType)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	tx.TransactionId = strconv.FormatInt(id, 10)

	return tx, nil
}
