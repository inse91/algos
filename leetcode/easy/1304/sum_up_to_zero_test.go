package _1304_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
	}{
		{
			name: "1",
			n:    5,
		},
		{
			name: "2",
			n:    3,
		},
		{
			name: "3",
			n:    1,
		},
		{
			name: "4",
			n:    2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := sumZero(c.n)
			var summ int
			for _, num := range res {
				summ += num
			}

			assert.Zero(t, summ)
		})
	}
}

func sumZero(n int) []int {
	res := make([]int, 0, n)
	if n%2 != 0 {
		res = append(res, 0)
	}

	for i := 1; i <= n/2; i++ {
		res = append(res, i, -i)
	}

	return res
}
