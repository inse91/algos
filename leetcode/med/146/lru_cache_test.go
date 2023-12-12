package _146

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type val struct {
	value int
	t     time.Time
}

type LRUCache struct {
	cap   int
	store map[int]val
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		store: make(map[int]val, capacity),
	}
}

func (lruc *LRUCache) Get(key int) int {
	if v, ok := lruc.store[key]; ok {
		lruc.store[key] = val{
			value: v.value,
			t:     time.Now(),
		}
		return v.value
	}
	return -1
}

func (lruc *LRUCache) Put(key int, value int) {
	now := time.Now()
	if len(lruc.store) < lruc.cap {
		lruc.store[key] = val{
			value: value,
			t:     time.Now(),
		}
		return
	}

	if _, ok := lruc.store[key]; ok {
		lruc.store[key] = val{
			value: value,
			t:     now,
		}
		return
	}

	t := now
	var keyToDelete int
	for k, v := range lruc.store {
		if v.t.Before(t) {
			keyToDelete = k
			t = v.t
		}
	}
	delete(lruc.store, keyToDelete)
	lruc.store[key] = val{
		value: value,
		t:     now,
	}
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRU(t *testing.T) {

	t.Run("1", func(t *testing.T) {

		lruCache := Constructor(2)

		lruCache.Put(1, 1)
		lruCache.Put(2, 2)
		v := lruCache.Get(1)
		assert.Equal(t, 1, v)

		lruCache.Put(3, 3)
		v = lruCache.Get(2)
		assert.Equal(t, -1, v)
		lruCache.Put(4, 4)
		v = lruCache.Get(1)
		assert.Equal(t, -1, v)

		v = lruCache.Get(3)
		assert.Equal(t, 3, v)
		v = lruCache.Get(4)
		assert.Equal(t, 4, v)
	})

	t.Run("2", func(t *testing.T) {

		lruCache := Constructor(2)

		v := lruCache.Get(2)
		assert.Equal(t, -1, v)

		lruCache.Put(2, 6)

		v = lruCache.Get(1)
		assert.Equal(t, -1, v)

		lruCache.Put(1, 5)
		lruCache.Put(1, 2)
		v = lruCache.Get(1)
		assert.Equal(t, 2, v)

		v = lruCache.Get(2)
		assert.Equal(t, 6, v)
	})
}
