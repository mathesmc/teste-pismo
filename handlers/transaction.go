package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	postgresql "github.com/mathesmc/teste-pismo/db/sqlc"
)

const (
	MISMATCH_OPERATION_TYPE    = "amount does not match with operation type"
	TRANSACTION_CREATION_ERROR = "error creating transaction"
)

type CreateTransactionsParams struct {
	AccountID       int64   `json:"account_id,omitempty"`
	OperationTypeID int64   `json:"operation_type_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
}

func (c *CreateTransactionsParams) Bind(r *http.Request) (err error) {

	if c.AccountID == 0 {
		err = ErrEmptyParam("account_id")
		return
	}

	if c.OperationTypeID == 0 {
		err = ErrEmptyParam("operation_type_id")
		return
	}

	if c.Amount == 0 {
		err = ErrEmptyParam("amount")
		return
	}

	return
}

// PostTrasaction create one account
//
//	@Summary	create one transaction with given
//	@Tags		accounts
//	@Param		transaction_params	body createTransactionsParams true "document number which we registrate transaction"
//	@Success	202				{object}	TransactionResponse
//	@Failure	404				{object}	ErrResponse
//	@Failure	422				{object}  ErrResponse
//	@Failure	500				{object}  ErrResponse
//	@Router		/transactions [post]
func PostTrasaction(w http.ResponseWriter, r *http.Request) {
	t := &CreateTransactionsParams{}

	if err := render.Bind(r, t); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, TRANSACTION_CREATION_ERROR))
		return
	}

	ot, err := postgresql.DB.GetOperationTypes(r.Context(), t.OperationTypeID)
	if err != nil {
		render.Render(w, r, &ErrResponse{Err: err, HTTPStatusCode: 500,
			StatusText: "database_error"})
	}

	if invalidAmount(ot, t.Amount) {
		err = ErrMismatchOperationType
		render.Render(w, r, ErrInvalidRequest(err, MISMATCH_OPERATION_TYPE))
		return
	}

	params := postgresql.CreateTransactionsParams{AccountID: t.AccountID, OperationTypeID: t.OperationTypeID, Amount: t.Amount}

	ct, err := postgresql.DB.CreateTransactions(context.Background(), params)

	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err, TRANSACTION_CREATION_ERROR))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewTransactionResponse(&ct))

}

type TransactionResponse struct {
	AccountID       int64   `json:"account_id"`
	OperationTypeID int64   `json:"operation_type"`
	Amount          float64 `json:"amount"`
}

func (a *TransactionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewTransactionResponse(a *postgresql.Transaction) *TransactionResponse {
	return &TransactionResponse{
		OperationTypeID: a.OperationTypeID,
		Amount:          a.Amount,
		AccountID:       a.AccountID,
	}
}

func invalidAmount(o postgresql.OperationType, a float64) bool {
	return o.Multiplier < 0 && a > 0 || o.Multiplier > 0 && a < 0
}
