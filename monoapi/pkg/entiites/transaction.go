package entiites

import "github.com/golang/protobuf/ptypes/timestamp"

type Transaction struct {
	ID       int
	UserID   int
	MonoInfo MonoTransaction
	Card Card
}

type MonoTransaction struct {
	ID              string
	Time            timestamp.Timestamp
	Description     string
	MCC             int
	Amount          int
	OperationAmount int
	CurrencyCode    int
	CommissionRate  int
	CashBackAmount  int
	Balance         int
	Hold            bool
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Password  string
	MonoToken string
	MonoUser MonoUser
}

type MonoUser struct {
	ID         string
	Name       string
	WebHookURL string
}

type Card struct {
	ID int
	UserID int
	Tracked bool
}

type MonoCard struct {
	ID           int
	CurrencyCode int
	CashbackType string
	Balance      int
	CreditLimit  int
	MaskedPan    []string
	Type         string
}
