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
		result, err := chk.Check(r)
		if err != nil {
			log.Printf("check failed: %s", err.Error())
		}

		switch result {
		case check.ACCEPT:
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
			return
		case check.REJECT:
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
