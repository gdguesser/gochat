package main

import (
	"github.com/bmizerany/pat"
	"github.com/gdguesser/gochat/internal/handlers"
	"net/http"
)

func routes() http.Handler  {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
