package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_statusHandler(t *testing.T) {
	t.Run("calling /api/status returns 200 OK", func(t *testing.T) {
		mux := New(nil)

		req := httptest.NewRequest("GET", "http://example.org/api/status", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		resp := w.Result()
		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "200 OK", resp.Status)
	})
}
