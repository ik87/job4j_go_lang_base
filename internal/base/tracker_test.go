package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tracker(t *testing.T) {

	t.Run("check link leak if get", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()

		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()

		res[0].Name = "Second Item"

		expected := []Item{
			{
				ID:   "1",
				Name: "First Item",
			},
		}

		assert.Equal(t,
			expected,
			tracker.GetItems(),
		)
	})

	t.Run("check link leak if add", func(t *testing.T) {
		t.Parallel()
		tracker := NewTracker()

		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		item.Name = "Second Item"

		expected := []Item{
			{
				ID:   "1",
				Name: "First Item",
			},
		}

		assert.Equal(t,
			expected,
			tracker.GetItems(),
		)
	})
}
