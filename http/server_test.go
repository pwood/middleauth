package http

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestServer_Start(t *testing.T) {
	t.Run("verify that the api router is attached to the http server", func(t *testing.T) {
		srv := &Server{Host: "127.0.0.1", Port: 0}

		done, err := srv.Start()
		assert.NoError(t, err)

		defer func() {
			assert.NoError(t, done(context.Background()))
		}()

		resp, err := http.Get(fmt.Sprintf("http://%s/api/status", srv.ln.Addr().String()))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
