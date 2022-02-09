package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pwood/middleauth/check"
	"github.com/pwood/middleauth/http/api"
	"log"
	"net"
	"net/http"
)

type Server struct {
	Host   string
	Port   int
	Checks []check.Checker
	ln     net.Listener
}

func (s *Server) Start() (func(context.Context) error, error) {
	apiRouter := api.New(s.Checks)

	r := mux.NewRouter()
	r.PathPrefix("/api").Handler(apiRouter)

	bindAddress := fmt.Sprintf("%s:%d", s.Host, s.Port)
	srv := &http.Server{Handler: r}

	ln, err := net.Listen("tcp", bindAddress)
	if err != nil {
		return nil, err
	}

	s.ln = ln

	go func() {
		if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("failed to list for http: %s", err.Error())
		}
	}()

	return func(ctx context.Context) error {
		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("couldn't shutdown server: %w", err)
		}

		return nil
	}, nil
}
