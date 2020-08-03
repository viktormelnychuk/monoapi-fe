package service

import (
	"context"

	"github.com/viktormelnychuk/monoapi/monoapi/pkg/entiites"
)

// MonoapiService describes the service.
type MonoapiService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Login(ctx context.Context, username, password string) (string, error)
	GetAllTransactions(ctx context.Context) ([]entiites.Transaction, error)
	GetTransaction(ctx context.Context, ID int) (entiites.Transaction, error)
	GetCards(ctx context.Context) ([]entiites.Card, error)
	EnableCard(ctx context.Context, cardId int, enabled bool) error
}

type basicMonoapiService struct{}

func (b *basicMonoapiService) Login(ctx context.Context, username string, password string) (string, error) {
	// TODO implement the business logic of Login
	return "e0",nil
}
func (b *basicMonoapiService) GetAllTransactions(ctx context.Context) (e0 []entiites.Transaction, e1 error) {
	// TODO implement the business logic of GetAllTransactions
	return e0, e1
}
func (b *basicMonoapiService) GetTransaction(ctx context.Context, ID int) (e0 entiites.Transaction, e1 error) {
	// TODO implement the business logic of GetTransaction
	return e0, e1
}
func (b *basicMonoapiService) GetCards(ctx context.Context) (e0 []entiites.Card, e1 error) {
	// TODO implement the business logic of GetCards
	return e0, e1
}
func (b *basicMonoapiService) EnableCard(ctx context.Context, cardId int, enabled bool) (e0 error) {
	// TODO implement the business logic of EnableCard
	return e0
}

// NewBasicMonoapiService returns a naive, stateless implementation of MonoapiService.
func NewBasicMonoapiService() MonoapiService {
	return &basicMonoapiService{}
}

// New returns a MonoapiService with all of the expected middleware wired in.
func New(middleware []Middleware) MonoapiService {
	var svc MonoapiService = NewBasicMonoapiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
