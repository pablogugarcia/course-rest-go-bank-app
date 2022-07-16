package domain

type Transaction struct {
	TransactionId   string
	TransactionType string
	Amount          float64
	AccountId       string
	OpeningDate     string
}
