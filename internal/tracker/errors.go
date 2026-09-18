package tracker

import "errors"

var ErrNotFound = errors.New("not found")

var ErrItemExists = errors.New("item with same id is exists")
