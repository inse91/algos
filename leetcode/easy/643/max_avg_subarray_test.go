package _643__test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxAvgSubarray(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  float64
	}{
		{
			name: "1",
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			res:  12.75,
		},
		{
			name: "2",
			nums: []int{5},
			k:    1,
			res:  5,
		},
		{
			name: "3",
			nums: []int{1, 2},
			k:    2,
			res:  1.5,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findMaxAverage(c.nums, c.k)
			assert.Equal(t, c.res, res)
		})
	}
}

func findMaxAverage(nums []int, k int) float64 {
	//slices.Sort(nums)
	lastIdx := len(nums)
	sl := nums[lastIdx-k : lastIdx]

	var sum1 float64
	for _, n := range sl {
		sum1 += float64(n)
	}

	for i := k; i < len(nums); i++ {
		sm := sum(nums[i-k : i])
		sum1 = max(sum1, float64(sm))
	}

	return sum1 / float64(k)
}

func sum(sl []int) (res int) {
	for i := range sl {
		res += sl[i]
	}

	return res
}
