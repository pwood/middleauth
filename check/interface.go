package check

import "net/http"

type Result uint

const (
	REJECT Result = iota
	NEXT
	ACCEPT
)

type Checker interface {
	Check(r *http.Request) (Result, error)
}
