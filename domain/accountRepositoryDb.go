package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/pablogugarcia/banking/errs"
	"github.com/pablogugarcia/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppErr) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func (d AccountRepositoryDb) FindById(accountId string) (Account, *errs.AppErr) {
	findQuery := "SELECT * FROM accounts where account_id = ?"

	var a Account

	err := d.client.Get(&a, findQuery, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return a, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while scaning customers table " + err.Error())
		return a, errs.NewUnexpectedError("unexpected database error")
	}

	return a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
