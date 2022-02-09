package static

import (
	"github.com/pwood/middleauth/check"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("initialising with an unknown result results in error", func(t *testing.T) {
		_, err := New(0xff)

		assert.Error(t, err)
	})

	t.Run("initialising with NEXT results in error", func(t *testing.T) {
		_, err := New(check.NEXT)

		assert.Error(t, err)
	})

	t.Run("initialising with ACCEPT returns static with no error", func(t *testing.T) {
		s, err := New(check.ACCEPT)

		assert.NoError(t, err)
		assert.Equal(t, check.ACCEPT, s.result)
	})
}

func TestStatic_Check(t *testing.T) {
	t.Run("returns initialised result type", func(t *testing.T) {
		expectedResult := check.ACCEPT

		s, err := New(expectedResult)
		assert.NoError(t, err)

		actualResult, err := s.Check(http.Request{})
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}
