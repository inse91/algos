package _706_test

import "testing"

func TestHashMap(t *testing.T) {

}

type MyHashMap struct {
	store map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		store: make(map[int]int),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.store[key] = value
}

func (this *MyHashMap) Get(key int) int {
	v, ok := this.store[key]
	if ok {
		return v
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.store, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
