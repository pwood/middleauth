package api

import (
	"github.com/pwood/middleauth/check"
	"log"
	"net/http"
	"strings"
)

type checkHandler struct {
	checks []check.Checker
}

func (c checkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decision := check.Decision{Result: check.Error}
	var err error

Checks:
	for _, chk := range c.checks {
		decision, err = chk.Check(r)
		if err != nil {
			log.Printf("check failed: %s", err.Error())
		}

		switch decision.Result {
		case check.Accept:
			break Checks
		case check.Reject:
			break Checks
		}
	}

	realIps := strings.Join(r.Header.Values("X-Real-Ip"), ",")
	log.Printf("%s: %s: %s (%s)", r.RemoteAddr, realIps, check.ResultNames[decision.Result], decision.Context)

	switch decision.Result {
	case check.Accept:
		http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
	case check.Reject:
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
