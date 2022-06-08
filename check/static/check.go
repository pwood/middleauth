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
	case check.Accept, check.Reject:
		return Static{result: result}, nil
	case check.Next:
		return Static{}, fmt.Errorf("Next result is invalid: %d", result)
	default:
		return Static{}, fmt.Errorf("unknown result set: %d", result)
	}
}

func (i Static) Check(_ *http.Request) (check.Decision, error) {
	return check.Decision{Result: i.result, Context: "Static"}, nil
}
