package static

import (
	"fmt"
	"github.com/pwood/middleauth/check"
	"net/http"
)

var _ check.Checker = (*Static)(nil)

type Static struct {
	result check.Result
}

func New(result check.Result) (Static, error) {
	switch result {
	case check.ACCEPT, check.REJECT:
		return Static{result: result}, nil
	case check.NEXT:
		return Static{}, fmt.Errorf("NEXT result is invalid: %d", result)
	default:
		return Static{}, fmt.Errorf("unknown result set: %d", result)
	}
}

func (i Static) Check(_ http.Request) (check.Result, error) {
	return i.result, nil
}
