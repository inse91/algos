package _705_test

import "testing"

func TestHashSet(t *testing.T) {

}

type MyHashSet struct {
	store map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		store: make(map[int]struct{}),
	}
}

func (this *MyHashSet) Add(key int) {
	this.store[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.store, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.store[key]
	return ok
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
