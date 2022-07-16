package service

import (
	"time"

	"github.com/pablogugarcia/banking/domain"
	dto "github.com/pablogugarcia/banking/dtos"
	"github.com/pablogugarcia/banking/errs"
	"github.com/pablogugarcia/banking/logger"
)

type TransactionService interface {
	NewTransaction(*dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr)
}
type DefaultTransactionService struct {
	repo           domain.TransactionRepository
	accountService domain.AccountRepository
}

func (s DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	tx := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		OpeningDate:     time.Now().Format("2006-01-02 15:04:05"),
		Amount:          req.Amount,
		TransactionType: req.Type,
	}
	acc, err := s.accountService.FindById(req.AccountId)
	if err != nil {
		logger.Error("Error in get account")
		return nil, err
	}
	if acc.Amount < req.Amount {
		return nil, errs.NewValidationError("The transactions amount is too high")
	}
	acc.Amount = acc.Amount - req.Amount
	_, err = s.accountService.Save(acc)
	if err != nil {
		logger.Error("Error in get account")
		return nil, err
	}
	transaction, err := s.repo.Save(&tx)
	if err != nil {
		return nil, err
	}

	return &dto.NewTransactionResponse{TransactionId: transaction.TransactionId, Balance: acc.Amount}, nil
}

func NewTransactionService(repo domain.TransactionRepository, accountService domain.AccountRepository) DefaultTransactionService {
	return DefaultTransactionService{repo, accountService}
}
