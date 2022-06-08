package api

import (
	"github.com/gorilla/mux"
	"github.com/pwood/middleauth/check"
	"net/http"
)

func New(checks []check.Checker) http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/api/status").HandlerFunc(statusHandler)
	r.PathPrefix("/api/check").Handler(checkHandler{checks: checks})

	return r
}
