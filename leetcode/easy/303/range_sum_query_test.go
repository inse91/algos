package _303_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeSumQuery(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		na := Constructor([]int{-2, 0, 3, -5, 2, -1})
		assert.Equal(t, 1, na.SumRange(0, 2))
		assert.Equal(t, -1, na.SumRange(2, 5))
		assert.Equal(t, -3, na.SumRange(0, 5))
	})
	t.Run("2", func(t *testing.T) {
		na := Constructor([]int{1})
		assert.Equal(t, 1, na.SumRange(0, 0))
	})

}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	prefixSums := make([]int, 0, len(nums))
	var acc int
	for _, num := range nums {
		prefixSums = append(prefixSums, acc+num)
		acc += num
	}
	return NumArray{nums: prefixSums}
}

func (na *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return na.nums[right]
	}
	return na.nums[right] - na.nums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
