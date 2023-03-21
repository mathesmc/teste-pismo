package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"

	"github.com/swaggo/http-swagger"

	_ "github.com/lib/pq"

	_ "github.com/mathesmc/teste-pismo/docs"
	"github.com/mathesmc/teste-pismo/handlers"
	"github.com/mathesmc/teste-pismo/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	ErrDocumentNumberIsEmpty       = fmt.Errorf("document_number should not be empty")
	ErrDocumentNumberDuplicatedKey = fmt.Errorf("document_number already registered in another account")
	connStr                        = "postgresql://root:secret@localhost:5432/mini_bank?sslmode=disable"
)

//@title Swagger Payment System
//@Version 1.0
//@description This project has symple operations for payments transactions
//@contact.name Matheus Morgado

//@host payments.swagger.io
//@BasePath /v2

func main() {

	err := storage.NewSqlRepo(connStr)
	if err != nil {
		syscall.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/accounts", func(r chi.Router) {
		r.Route("/{account_id}", func(r chi.Router) {
			r.Use(handlers.AccountCtx)
			r.Get("/", handlers.GetAccount)
		})
		r.Post("/", handlers.PostAccount)

	})

	r.Post("/transactions", handlers.PostTrasaction)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/index/doc.json"), //The url pointing to API definition
	))

	log.Println("server starting on port 3000")
	http.ListenAndServe(":3000", r)
}
