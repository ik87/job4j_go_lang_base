package tracker

import (
	"fmt"
	"slices"
)

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) AddItem(item Item) {
	t.items = append(t.items, item)
}

func (t *Tracker) DelItem(uuid string) {
	t.items = slices.DeleteFunc(t.items, func(it Item) bool {
		return it.ID == uuid
	})
}

func (t *Tracker) UpdateItem(item Item) {
	for i := 0; i < len(t.items); i++ {
		if item.ID == t.items[i].ID {
			t.items[i].Name = item.Name
			return
		}
	}
}
