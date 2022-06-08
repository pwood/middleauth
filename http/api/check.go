package api

import (
	"github.com/pwood/middleauth/check"
	"log"
	"net/http"
)

type checkHandler struct {
	checks []check.Checker
}

func (c checkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, chk := range c.checks {
		decision, err := chk.Check(r)
		if err != nil {
			log.Printf("check failed: %s", err.Error())
		}

		switch decision.Result {
		case check.Accept:
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
			return
		case check.Reject:
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
