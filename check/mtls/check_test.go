package mtls

import (
	"github.com/pwood/middleauth/check"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("returns a new MTLS with desired result on match", func(t *testing.T) {
		out, err := New(check.Accept)

		assert.NoError(t, err)
		assert.Equal(t, check.Accept, out.result)
	})
}

func TestMTLS_Check(t *testing.T) {
	t.Run("returns initialised result type if header is present", func(t *testing.T) {
		expectedResult := check.Accept

		s, err := New(expectedResult)
		assert.NoError(t, err)

		r, err := http.NewRequest(http.MethodGet, "", nil)
		assert.NoError(t, err)

		r.Header.Set(header, "present")

		decision, err := s.Check(r)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, decision.Result)
	})

	t.Run("returns next result type if header is not present", func(t *testing.T) {
		expectedResult := check.Next

		s, err := New(check.Accept)
		assert.NoError(t, err)

		r, err := http.NewRequest(http.MethodGet, "", nil)
		assert.NoError(t, err)

		decision, err := s.Check(r)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, decision.Result)
	})
}
