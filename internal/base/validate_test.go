package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("res - [\"req.UserID is empty\", \"req.Title is empty\", \"req.Description is empty\"]", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{UserID: "", Title: "", Description: ""}
		rsl := base.Validate(&in)

		assert.Equal(t, []string{"req.UserID is empty", "req.Title is empty", "req.Description is empty"}, rsl)
	})

	t.Run("res - [\"req.Title is empty\", \"req.Description is empty\"]", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{UserID: "1", Title: "", Description: ""}
		rsl := base.Validate(&in)

		assert.Equal(t, []string{"req.Title is empty", "req.Description is empty"}, rsl)
	})

	t.Run("res - []", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{UserID: "1", Title: "Hello", Description: "Hello World"}
		rsl := base.Validate(&in)

		assert.Equal(t, []string{}, rsl)
	})

	t.Run("res - [req is nil]", func(t *testing.T) {
		t.Parallel()

		var in *base.ValidateRequest
		rsl := base.Validate(in)

		assert.Equal(t, []string{"req is nil"}, rsl)
	})
}
