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

func (t *Tracker) AddItem(item Item) (Item, error) {
	_, ok := t.indexOf(item.ID)
	if ok {
		return item, fmt.Errorf("%w: %s", ErrItemExists, item.ID)
	}
	t.items = append(t.items, item)
	return item, nil
}

func (t *Tracker) DelItem(uuid string) {
	t.items = slices.DeleteFunc(t.items, func(it Item) bool {
		return it.ID == uuid
	})
}

func (t *Tracker) UpdateItem(item Item) error {
	i, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}
	t.items[i].Name = item.Name
	return nil

}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
}
