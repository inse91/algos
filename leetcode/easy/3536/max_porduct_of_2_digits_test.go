package _1636_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-product-of-two-digits/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    31,
			res:  3,
		},
		{
			name: "2",
			n:    22,
			res:  4,
		},
		{
			name: "3",
			n:    124,
			res:  8,
		},
		{
			name: "44",
			n:    712585593,
			res:  72,
		},
		{
			name: "47",
			n:    989,
			res:  81,
		},
		{
			name: "52",
			n:    543,
			res:  20,
		},
		{
			name: "57",
			n:    5455553,
			res:  25,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxProduct(c.n))
		})
	}
}

func maxProduct(n int) int {
	if n < 100 {
		first := n / 10
		second := n - first*10
		return first * second
	}

	digs := [10]int{}
	for n > 0 {
		digs[n%10] += 1
		n /= 10
	}

	var c int
	res := 1
	for i := 9; i >= 0; i-- {
		for ; digs[i] > 0 && c != 2; c++ {
			res *= i
			digs[i] -= 1
		}
	}

	return res
}
