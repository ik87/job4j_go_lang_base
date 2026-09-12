package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Mono(t *testing.T) {
	t.Parallel()

	t.Run("{first:cat, second:dog, third:bird} AND size = 3 - {first:cat, second:dog, third:bird}", func(t *testing.T) {
		t.Parallel()

		cache := NewLruCache(3)
		cache.Put("first", "cat")
		cache.Put("second", "dog")
		cache.Put("third", "bird")

		assert.Equal(t, "cat", *cache.Get("first"))
		assert.Equal(t, "dog", *cache.Get("second"))
		assert.Equal(t, "bird", *cache.Get("third"))
		assert.Nil(t, cache.Get("elephant"))
	})

	t.Run("{first:cat, second:dog, third:bird} AND size = 2 - {second:dog, third:bird}", func(t *testing.T) {
		t.Parallel()

		cache := NewLruCache(2)
		cache.Put("first", "cat")
		cache.Put("second", "dog")
		cache.Put("third", "bird")

		assert.Equal(t, "dog", *cache.Get("second"))
		assert.Equal(t, "bird", *cache.Get("third"))
		assert.Nil(t, cache.Get("first"))
	})
}
