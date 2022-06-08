package check

import "net/http"

type Result uint

const (
	Reject Result = iota
	Next
	Accept
)

type Decision struct {
	Result  Result
	Context string
}

var NextDecision = Decision{Result: Next}

type Checker interface {
	Check(r *http.Request) (Decision, error)
}
