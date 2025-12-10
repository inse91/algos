package _1317_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  []int
	}{
		{
			name: "1",
			num:  2,
			res:  []int{1, 1},
		},
		{
			name: "2",
			num:  11,
			res:  []int{2, 9},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, getNoZeroIntegers(c.num))
		})
	}
}

func getNoZeroIntegers(n1 int) []int {
	var n2 int
	for c := 1; n1 > 0; c++ {
		n1 -= 1
		n2 += 1
		if containsZero(n1) {
			continue
		}
		if containsZero(n2) {
			continue
		}

		return []int{n1, n2}
	}

	return nil
}

func containsZero(n int) bool {
	s := strconv.Itoa(n)
	for _, r := range s {
		if r == '0' {
			return true
		}
	}

	return false
}
