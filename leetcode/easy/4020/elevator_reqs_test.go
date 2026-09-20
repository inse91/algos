package _4020_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/elevator-requests-i/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		reqs []int
		res  int
	}{
		{
			name: "1",
			reqs: []int{2, 1, 4, 3},
			n:    5,
			res:  7,
		},
		{
			name: "2",
			reqs: []int{2, 0, 0},
			n:    3,
			res:  4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, elevatorRequests(c.n, c.reqs))
		})
	}
}

func elevatorRequests(n int, requests []int) int {
	var (
		cur int
		res int
	)

	for _, v := range requests {
		res += abs(v - cur)
		cur = v
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
