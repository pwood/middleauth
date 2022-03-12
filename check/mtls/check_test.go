package mtls

import (
	"github.com/pwood/middleauth/check"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("returns a new MTLS with desired result on match", func(t *testing.T) {
		out, err := New(check.ACCEPT)

		assert.NoError(t, err)
		assert.Equal(t, check.ACCEPT, out.result)
	})
}

func TestMTLS_Check(t *testing.T) {
	t.Run("returns initialised result type if header is present", func(t *testing.T) {
		expectedResult := check.ACCEPT

		s, err := New(expectedResult)
		assert.NoError(t, err)

		r, err := http.NewRequest(http.MethodGet, "", nil)
		assert.NoError(t, err)

		r.Header.Set(header, "present")

		actualResult, err := s.Check(r)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("returns next result type if header is not present", func(t *testing.T) {
		expectedResult := check.NEXT

		s, err := New(check.ACCEPT)
		assert.NoError(t, err)

		r, err := http.NewRequest(http.MethodGet, "", nil)
		assert.NoError(t, err)

		actualResult, err := s.Check(r)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}
