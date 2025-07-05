package _2094_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/finding-3-digit-even-numbers

func Test(t *testing.T) {
	cases := []struct {
		name   string
		digits []int
		res    []int
	}{
		{
			name:   "1",
			digits: []int{2, 1, 3, 0},
			res:    []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
		},
		{
			name:   "2",
			digits: []int{2, 2, 8, 8, 2},
			res:    []int{222, 228, 282, 288, 822, 828, 882},
		},
		{
			name:   "3",
			digits: []int{3, 7, 5},
		},
		{
			name:   "53",
			digits: []int{2, 2, 8, 8, 2, 3, 1, 7},
			res: []int{
				122, 128, 132, 138, 172, 178, 182, 188, 212, 218, 222, 228,
				232, 238, 272, 278, 282, 288, 312, 318, 322, 328, 372, 378,
				382, 388, 712, 718, 722, 728, 732, 738, 782, 788, 812, 818,
				822, 828, 832, 838, 872, 878, 882,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, findEvenNumbers(c.digits))
		})
	}
}
func findEvenNumbers(digits []int) []int {
	if len(digits) < 3 {
		return nil
	}

	var evens, c int
	var digsCounts [10]int
	for _, d := range digits {
		if d%2 == 0 {
			evens++
		}
		if digsCounts[d] == 0 {
			c++
		}
		digsCounts[d]++
	}

	if evens == 0 {
		return nil
	}

	res := make([]int, 0, c*evens)
	for i := 100; i <= 998; i += 2 {
		var cur [10]int
		for j, count := range digsCounts {
			cur[j] = count
		}

		num := i
		var d int
		var innerCycleBroken bool
		for num > 0 {
			d = num % 10
			num /= 10
			if cur[d] == 0 {
				innerCycleBroken = true
				break
			}

			cur[d]--
		}

		if innerCycleBroken {
			continue
		}

		res = append(res, i)
	}

	return res
}
