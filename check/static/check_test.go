package static

import (
	"github.com/pwood/middleauth/check"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("initialising with an unknown result results in error", func(t *testing.T) {
		_, err := New(0xff)

		assert.Error(t, err)
	})

	t.Run("initialising with Next results in error", func(t *testing.T) {
		_, err := New(check.Next)

		assert.Error(t, err)
	})

	t.Run("initialising with Accept returns static with no error", func(t *testing.T) {
		s, err := New(check.Accept)

		assert.NoError(t, err)
		assert.Equal(t, check.Accept, s.result)
	})
}

func TestStatic_Check(t *testing.T) {
	t.Run("returns initialised result type", func(t *testing.T) {
		expectedResult := check.Accept

		s, err := New(expectedResult)
		assert.NoError(t, err)

		decision, err := s.Check(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, decision.Result)
	})
}
