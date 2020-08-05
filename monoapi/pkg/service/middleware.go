package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	entiites "github.com/viktormelnychuk/monoapi/monoapi/pkg/entiites"
)

// Middleware describes a service middleware.
type Middleware func(MonoapiService) MonoapiService

type loggingMiddleware struct {
	logger log.Logger
	next   MonoapiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a MonoapiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next MonoapiService) MonoapiService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Login(ctx context.Context, username string, password string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "Login", "username", username, "password", password, "s0", s0, "e1", e1)
	}()
	return l.next.Login(ctx, username, password)
}
func (l loggingMiddleware) SignUp(ctx context.Context, user entiites.User) (e0 error) {
	defer func() {
		l.logger.Log("method", "SignUp", "user", user, "e0", e0)
	}()
	return l.next.SignUp(ctx, user)
}
func (l loggingMiddleware) GetAllTransactions(ctx context.Context) (e0 []entiites.Transaction, e1 error) {
	defer func() {
		l.logger.Log("method", "GetAllTransactions", "e0", e0, "e1", e1)
	}()
	return l.next.GetAllTransactions(ctx)
}
func (l loggingMiddleware) GetTransaction(ctx context.Context, ID int) (e0 entiites.Transaction, e1 error) {
	defer func() {
		l.logger.Log("method", "GetTransaction", "ID", ID, "e0", e0, "e1", e1)
	}()
	return l.next.GetTransaction(ctx, ID)
}
func (l loggingMiddleware) GetCards(ctx context.Context) (e0 []entiites.Card, e1 error) {
	defer func() {
		l.logger.Log("method", "GetCards", "e0", e0, "e1", e1)
	}()
	return l.next.GetCards(ctx)
}
func (l loggingMiddleware) EnableCard(ctx context.Context, cardId int, enabled bool) (e0 error) {
	defer func() {
		l.logger.Log("method", "EnableCard", "cardId", cardId, "enabled", enabled, "e0", e0)
	}()
	return l.next.EnableCard(ctx, cardId, enabled)
}
