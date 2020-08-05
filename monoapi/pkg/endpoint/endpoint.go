package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	entiites "github.com/viktormelnychuk/monoapi/monoapi/pkg/entiites"
	service "github.com/viktormelnychuk/monoapi/monoapi/pkg/service"
)

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.MonoapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		s0, e1 := s.Login(ctx, req.Username, req.Password)
		return LoginResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.E1
}

// SignUpRequest collects the request parameters for the SignUp method.
type SignUpRequest struct {
	User entiites.User `json:"user"`
}

// SignUpResponse collects the response parameters for the SignUp method.
type SignUpResponse struct {
	E0 error `json:"e0"`
}

// MakeSignUpEndpoint returns an endpoint that invokes SignUp on the service.
func MakeSignUpEndpoint(s service.MonoapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		e0 := s.SignUp(ctx, req.User)
		return SignUpResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SignUpResponse) Failed() error {
	return r.E0
}

// GetAllTransactionsRequest collects the request parameters for the GetAllTransactions method.
type GetAllTransactionsRequest struct{}

// GetAllTransactionsResponse collects the response parameters for the GetAllTransactions method.
type GetAllTransactionsResponse struct {
	E0 []entiites.Transaction `json:"e0"`
	E1 error                  `json:"e1"`
}

// MakeGetAllTransactionsEndpoint returns an endpoint that invokes GetAllTransactions on the service.
func MakeGetAllTransactionsEndpoint(s service.MonoapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		e0, e1 := s.GetAllTransactions(ctx)
		return GetAllTransactionsResponse{
			E0: e0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAllTransactionsResponse) Failed() error {
	return r.E1
}

// GetTransactionRequest collects the request parameters for the GetTransaction method.
type GetTransactionRequest struct {
	ID int `json:"id"`
}

// GetTransactionResponse collects the response parameters for the GetTransaction method.
type GetTransactionResponse struct {
	E0 entiites.Transaction `json:"e0"`
	E1 error                `json:"e1"`
}

// MakeGetTransactionEndpoint returns an endpoint that invokes GetTransaction on the service.
func MakeGetTransactionEndpoint(s service.MonoapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTransactionRequest)
		e0, e1 := s.GetTransaction(ctx, req.ID)
		return GetTransactionResponse{
			E0: e0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTransactionResponse) Failed() error {
	return r.E1
}

// GetCardsRequest collects the request parameters for the GetCards method.
type GetCardsRequest struct{}

// GetCardsResponse collects the response parameters for the GetCards method.
type GetCardsResponse struct {
	E0 []entiites.Card `json:"e0"`
	E1 error           `json:"e1"`
}

// MakeGetCardsEndpoint returns an endpoint that invokes GetCards on the service.
func MakeGetCardsEndpoint(s service.MonoapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		e0, e1 := s.GetCards(ctx)
		return GetCardsResponse{
			E0: e0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetCardsResponse) Failed() error {
	return r.E1
}

// EnableCardRequest collects the request parameters for the EnableCard method.
type EnableCardRequest struct {
	CardId  int  `json:"card_id"`
	Enabled bool `json:"enabled"`
}

// EnableCardResponse collects the response parameters for the EnableCard method.
type EnableCardResponse struct {
	E0 error `json:"e0"`
}

// MakeEnableCardEndpoint returns an endpoint that invokes EnableCard on the service.
func MakeEnableCardEndpoint(s service.MonoapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EnableCardRequest)
		e0 := s.EnableCard(ctx, req.CardId, req.Enabled)
		return EnableCardResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r EnableCardResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, username string, password string) (s0 string, e1 error) {
	request := LoginRequest{
		Password: password,
		Username: username,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).S0, response.(LoginResponse).E1
}

// SignUp implements Service. Primarily useful in a client.
func (e Endpoints) SignUp(ctx context.Context, user entiites.User) (e0 error) {
	request := SignUpRequest{User: user}
	response, err := e.SignUpEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SignUpResponse).E0
}

// GetAllTransactions implements Service. Primarily useful in a client.
func (e Endpoints) GetAllTransactions(ctx context.Context) (e0 []entiites.Transaction, e1 error) {
	request := GetAllTransactionsRequest{}
	response, err := e.GetAllTransactionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllTransactionsResponse).E0, response.(GetAllTransactionsResponse).E1
}

// GetTransaction implements Service. Primarily useful in a client.
func (e Endpoints) GetTransaction(ctx context.Context, ID int) (e0 entiites.Transaction, e1 error) {
	request := GetTransactionRequest{ID: ID}
	response, err := e.GetTransactionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTransactionResponse).E0, response.(GetTransactionResponse).E1
}

// GetCards implements Service. Primarily useful in a client.
func (e Endpoints) GetCards(ctx context.Context) (e0 []entiites.Card, e1 error) {
	request := GetCardsRequest{}
	response, err := e.GetCardsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetCardsResponse).E0, response.(GetCardsResponse).E1
}

// EnableCard implements Service. Primarily useful in a client.
func (e Endpoints) EnableCard(ctx context.Context, cardId int, enabled bool) (e0 error) {
	request := EnableCardRequest{
		CardId:  cardId,
		Enabled: enabled,
	}
	response, err := e.EnableCardEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(EnableCardResponse).E0
}
