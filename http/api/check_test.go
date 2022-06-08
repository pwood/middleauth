package api

import (
	"github.com/pwood/middleauth/check"
	"github.com/pwood/middleauth/check/static"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_checkHandler_ServeHTTP(t *testing.T) {
	t.Run("calling /api/check with a single checker that ACCEPTs the request returns 200 OK", func(t *testing.T) {
		s, err := static.New(check.Accept)
		assert.NoError(t, err)

		mux := New([]check.Checker{s})

		req := httptest.NewRequest("GET", "http://example.org/api/check", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		resp := w.Result()
		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "200 OK", resp.Status)
	})

	t.Run("calling /api/check with a single checker that REJECTs the request returns 200 OK", func(t *testing.T) {
		s, err := static.New(check.Reject)
		assert.NoError(t, err)

		mux := New([]check.Checker{s})

		req := httptest.NewRequest("GET", "http://example.org/api/check", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		resp := w.Result()
		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		assert.Equal(t, "403 Forbidden", resp.Status)
	})

	t.Run("calling /api/check with no checks returns 500 Internal Server Error", func(t *testing.T) {
		mux := New([]check.Checker{})

		req := httptest.NewRequest("GET", "http://example.org/api/check", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		resp := w.Result()
		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, "500 Internal Server Error", resp.Status)
	})
}
