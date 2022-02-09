package http

import "github.com/pwood/middleauth/check"

type Server struct {
	Port   int
	Checks []check.Checker
}

func (s Server) Start() (func() error, error) {
	panic("unimplemented")
}
