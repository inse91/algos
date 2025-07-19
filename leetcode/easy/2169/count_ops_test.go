package _2169_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-operations-to-obtain-zero

func Test(t *testing.T) {
	cases := []struct {
		name       string
		num1, num2 int
		res        int
	}{
		{
			name: "1",
			num1: 2,
			num2: 3,
			res:  3,
		},
		{
			name: "2",
			num1: 10,
			num2: 10,
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countOperations(c.num1, c.num2))
		})
	}
}

func countOperations(num1, num2 int) int {
	var c int
	for {
		if num1 == 0 || num2 == 0 {
			return c
		}

		c++
		if num1 <= num2 {
			num2 -= num1
			continue
		}

		num1 -= num2
	}
}
