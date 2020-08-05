package entiites

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Transaction struct {
	ID       int
	UserID   int
	MonoInfo MonoTransaction
	Card     Card
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
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	MonoToken string `json:"mono_token"`
	MonoUser  MonoUser
}

func (u *User) String() string {
	return fmt.Sprintf("username: %s, first_name: %s, last_name: %s", u.UserName, u.FirstName, u.LastName)
}

type MonoUser struct {
	ID         string
	Name       string
	WebHookURL string
}

type Card struct {
	ID      int
	UserID  int
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
