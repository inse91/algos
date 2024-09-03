package _1342_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  14,
			res:  6,
		},
		{
			name: "2",
			num:  8,
			res:  4,
		},
		{
			name: "3",
			num:  123,
			res:  12,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numberOfSteps(c.num))
		})
	}
}

func numberOfSteps(num int) int {
	var res int
	for num != 0 {
		res++

		if num%2 == 0 {
			num /= 2
			continue
		}

		num -= 1
	}

	return res
}
