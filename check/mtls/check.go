package mtls

import (
	"github.com/pwood/middleauth/check"
	"net/http"
)

var _ check.Checker = (*MTLS)(nil)

type MTLS struct {
	result check.Result
}

func New(result check.Result) (MTLS, error) {
	m := MTLS{result: result}

	return m, nil
}

const header string = "X-Forwarded-Tls-Client-Cert"

func (m MTLS) Check(r *http.Request) (check.Decision, error) {
	if _, found := r.Header[header]; found {
		return check.Decision{Result: m.result, Context: "MTLS"}, nil
	}

	return check.NextDecision, nil
}
