package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Errors(t *testing.T) {
	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker.UpdateItem(item)
		assert.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("error add - item is exists", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		var err error

		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		_, err = tracker.AddItem(item)

		assert.Nil(t, err)

		item2 := Item{
			ID:   "1",
			Name: "Second Item",
		}

		_, err = tracker.AddItem(item2)

		assert.ErrorIs(t, err, ErrItemExists)
	})
}
