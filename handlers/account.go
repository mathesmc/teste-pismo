package handlers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	postgresql "github.com/mathesmc/teste-pismo/db/sqlc"
	storage "github.com/mathesmc/teste-pismo/storage"
	_ "github.com/swaggo/http-swagger"
)

const (
	DOCUMENT_NUMBER         = "document_number"
	INVALID_DOCUMENT_LENGTH = "documen_number is allowed to have only 11 characters"
	CREATING_ACCOUNT_ERROR  = "error creating account"
	DUPLICATE_KEY           = "duplicate key"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type GetAccountResponse struct {
	DocumentNumber string `json:"document_number,omitempty"`
	ID             int64  `json:"id,omitempty"`
}

type CreatedAccountResponse struct {
	DocumentNumber string `json:"document_number,omitempty"`
}

// GetAccount return one account
//
//	@Summary	get one account
//	@Tags		accounts
//	@Prouce		json
//	@Success	200 {object}		getAccountResponse
//	@Failure	404	{object}    ErrResponse
//	@Failure	500	{object}    ErrResponse
//	@Router		/accounts/{id} [get]
func GetAccount(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(postgresql.Account)
	if err := render.Render(w, r, showAccountResponse(&account)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func PostAccountResponse(a *postgresql.Account) *CreatedAccountResponse {
	return &CreatedAccountResponse{DocumentNumber: a.DocumentNumber}

}

func (p *CreatedAccountResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func showAccountResponse(a *postgresql.Account) *GetAccountResponse {
	return &GetAccountResponse{DocumentNumber: a.DocumentNumber, ID: a.ID}

}

func (a *GetAccountResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func AccountCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			a          postgresql.Account
			err        error
			account_id = "account_id"
		)

		if accountID := chi.URLParam(r, account_id); accountID != "" {
			id, _ := strconv.ParseInt(accountID, 10, 64)
			a, err = storage.Repo.GetAccount(id)
		}
		if err != nil || a.ID == 0 {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "account", a)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// PostAccount create one account
//
//	@Summary	create one account with given document_number
//	@Tags		accounts
//	@Param		document_number	body createAccountParams true "represents the document number registered in one account"
//	@Success	202				{object}	createdAccountResponse
//	@Failure	404				{object}	ErrResponse
//	@Failure	422				{object}  ErrResponse
//	@Failure	422				{object}  ErrResponse
//	@Failure	500				{object}  ErrResponse
//	@Router		/accounts [post]
func PostAccount(w http.ResponseWriter, r *http.Request) {
	a := &CreateAccountParams{}
	if err := render.Bind(r, a); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, CREATING_ACCOUNT_ERROR))
		return
	}

	cAccount, err := postgresql.DB.CreateAccount(context.Background(), a.DocumentNumber)
	if err != nil {
		if strings.Contains(err.Error(), DUPLICATE_KEY) {
			err = ErrDuplicatedKey(DOCUMENT_NUMBER)
		}

		render.Render(w, r, ErrInvalidRequest(err, CREATING_ACCOUNT_ERROR))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, PostAccountResponse(&cAccount))
	return
}

func (p Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type CreateAccountParams struct {
	DocumentNumber string `json:"document_number,omitempty"`
}

func (c *CreateAccountParams) Bind(r *http.Request) error {

	switch {
	case c.DocumentNumber == "":
		return ErrEmptyParam(DOCUMENT_NUMBER)
	case len(c.DocumentNumber) != 11:
		return ErrInvalidDocument(INVALID_DOCUMENT_LENGTH)
	default:
		return nil
	}

}
