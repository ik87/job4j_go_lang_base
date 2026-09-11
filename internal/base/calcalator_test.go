package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Add(t *testing.T) {
	t.Parallel()

	t.Run("1 + 2 = 3", func(t *testing.T) {
		t.Parallel()
		rsl := base.Add(1, 2)
		expected := 3

		assert.Equal(t, rsl, expected)
	})
}
