package _485_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConseqOnes(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 1, 0, 1, 1, 1},
			res:  3,
		},
		{
			name: "2",
			nums: []int{1, 0, 1, 1, 0, 1},
			res:  2,
		},
		{
			name: "3",
			nums: []int{0, 0, 0, 0, 0, 0},
			res:  0,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findMaxConsecutiveOnes(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func findMaxConsecutiveOnes(nums []int) int {
	var cnt int
	var maxCnt int
	for _, num := range nums {
		if num == 1 {
			cnt++
			continue
		}
		maxCnt = max(maxCnt, cnt)
		cnt = 0
	}

	return max(maxCnt, cnt)
}
