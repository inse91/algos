package _2042_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "1 box has 3 blue 4 red 6 green and 12 yellow marbles",
			res:  true,
		},
		{
			name: "1_1",
			s:    "1 box has 13 blue 4 red 6 green and 12 yellow marbles",
			res:  false,
		},
		{
			name: "1_2",
			s:    "1 box has 3blue4 red 6green and12yellow marbles",
			res:  true,
		},
		{
			name: "2",
			s:    "hello 5 x 5",
			res:  false,
		},
		{
			name: "2_2",
			s:    "hello 5x5",
			res:  false,
		},
		{
			name: "3",
			s:    "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, areNumbersAscending(c.s))
		})
	}
}

func areNumbersAscending(s string) bool {
	var (
		prev    int
		current int
		nums    []int
		reset   bool
	)
	for _, r := range s {
		if unicode.IsDigit(r) {
			reset = true
			num := int(r - '0')
			nums = append(nums, num)

			continue
		}

		if !reset {
			continue
		}

		current = getSum(nums)
		if current <= prev {
			return false
		}
		nums = nums[:0]
		prev = current
		reset = false
	}

	if len(nums) != 0 {
		current = getSum(nums)
		if current <= prev {
			return false
		}
	}

	return true
}

func getSum(digits []int) int {
	var res int
	for i, d := range digits {
		for range len(digits) - i - 1 {
			d *= 10
		}
		res += d
	}

	return res
}
