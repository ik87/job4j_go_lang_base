package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Mono(t *testing.T) {
	t.Parallel()

	t.Run("[1, 2, 3] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[1, 1, 1] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 1, 1}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[3, 2, 1] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{3, 2, 1}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[3, 2, 4] - false", func(t *testing.T) {
		t.Parallel()

		in := []int{3, 2, 4}
		rsl := base.Mono(in)

		assert.Equal(t, false, rsl)
	})

	t.Run("[1, 10, 100, 0] - false", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 10, 100, 0}
		rsl := base.Mono(in)

		assert.Equal(t, false, rsl)
	})

	t.Run("[1, 2, 2, 3, 4] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 2, 3, 4}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[10, 8, 6, 4, 2] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{10, 8, 6, 4, 2}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[10, 8, 8, 5, 1] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{10, 8, 8, 5, 1}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[5, 4, 4, 6] - false", func(t *testing.T) {
		t.Parallel()

		in := []int{5, 4, 4, 6}
		rsl := base.Mono(in)

		assert.Equal(t, false, rsl)
	})

	t.Run("[-1, -2, -3, -4] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{-1, -2, -3, -4}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[-5, 0, 3, 10] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{-5, 0, 3, 10}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[0, -1, 2] - false", func(t *testing.T) {
		t.Parallel()

		in := []int{0, -1, 2}
		rsl := base.Mono(in)

		assert.Equal(t, false, rsl)
	})

	t.Run("[0, 0, 1, 1, 2, 2] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{0, 0, 1, 1, 2, 2}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})

	t.Run("[0, 0, -1, -1, -2, -2] - true", func(t *testing.T) {
		t.Parallel()

		in := []int{0, 0, -1, -1, -2, -2}
		rsl := base.Mono(in)

		assert.Equal(t, true, rsl)
	})
}
