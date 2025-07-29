package _2600_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/k-items-with-the-maximum-sum

func Test(t *testing.T) {
	cases := []struct {
		name   string
		ones   int
		zeroes int
		negs   int
		k      int
		res    int
	}{
		{
			name:   "1",
			ones:   3,
			zeroes: 2,
			negs:   0,
			k:      2,
			res:    2,
		},
		{
			name:   "2",
			ones:   3,
			zeroes: 2,
			negs:   0,
			k:      4,
			res:    3,
		},
		{
			name:   "3",
			ones:   3,
			zeroes: 2,
			negs:   1,
			k:      6,
			res:    2,
		},
		{
			name:   "4",
			ones:   3,
			zeroes: 2,
			negs:   1,
			k:      5,
			res:    3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, kItemsWithMaximumSum(c.ones, c.zeroes, c.negs, c.k))
		})
	}
}

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	var sum int

	sum += min(numOnes, k)
	k -= sum + numZeros
	if k <= 0 {
		return sum
	}

	sum -= min(numNegOnes, k)

	return sum
}
