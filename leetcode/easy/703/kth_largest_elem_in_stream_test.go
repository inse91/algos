package _703_test

// https://leetcode.com/problems/kth-largest-element-in-a-stream/description/

import (
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

type KthLargest struct {
	k     int
	cur   int
	store []int
}

func Constructor(k int, nums []int) KthLargest {
	slices.SortFunc(nums, func(a, b int) int {
		if a < b {
			return 1
		}
		return -1
	})

	if k <= len(nums) {
		nums = nums[:k]
	}

	ln := len(nums)
	for i := 0; ln+i < k; i++ {
		nums = append(nums, -10_001)
	}

	return KthLargest{
		k:     k,
		store: nums,
		cur:   nums[k-1],
	}
}

func (this *KthLargest) Add(val int) int {
	if val <= this.cur {
		return this.cur
	}

	if len(this.store) == 0 {
		this.store = append(this.store, val)
		this.cur = val

		return val
	}

	// TODO binary idx search
	idx := -1
	for i, v := range this.store {
		if val >= v {
			idx = i
			break
		}
	}

	this.store = slices.Insert(this.store[:len(this.store)-1], idx, val)
	this.cur = this.store[len(this.store)-1]

	return this.cur
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func Test(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		kle := Constructor(3, []int{4, 5, 8, 2})

		require.Equal(t, 4, kle.Add(3))
		require.Equal(t, 5, kle.Add(5))
		require.Equal(t, 5, kle.Add(10))
		require.Equal(t, 8, kle.Add(9))
		require.Equal(t, 8, kle.Add(4))
	})

	t.Run("1", func(t *testing.T) {
		kle := Constructor(1, []int{})

		require.Equal(t, -3, kle.Add(-3))
		require.Equal(t, -2, kle.Add(-2))
		require.Equal(t, -2, kle.Add(-4))
		require.Equal(t, 0, kle.Add(0))
		require.Equal(t, 4, kle.Add(4))
	})

	t.Run("6", func(t *testing.T) {
		kle := Constructor(1, []int{-2})

		require.Equal(t, -2, kle.Add(-3))
		require.Equal(t, 0, kle.Add(0))
		require.Equal(t, 2, kle.Add(2))
		require.Equal(t, 2, kle.Add(-1))
		require.Equal(t, 4, kle.Add(4))
	})

	t.Run("7", func(t *testing.T) {
		kle := Constructor(2, []int{0})

		require.Equal(t, -1, kle.Add(-1))
		require.Equal(t, 0, kle.Add(1))
		require.Equal(t, 0, kle.Add(-2))
		require.Equal(t, 0, kle.Add(-4))
		require.Equal(t, 1, kle.Add(3))
	})

	t.Run("9", func(t *testing.T) {
		kle := Constructor(3, []int{5, -1})

		require.Equal(t, -1, kle.Add(2))
		require.Equal(t, 1, kle.Add(1))
		require.Equal(t, 1, kle.Add(-1))
		require.Equal(t, 2, kle.Add(3))
		require.Equal(t, 3, kle.Add(4))
	})
}
