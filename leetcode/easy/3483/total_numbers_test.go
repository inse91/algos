package _3483_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/unique-3-digit-even-numbers

func Test(t *testing.T) {
	cases := []struct {
		name   string
		digits []int
		k      int
		res    int
	}{
		{
			name:   "1",
			digits: []int{1, 2, 3, 4},
			res:    12,
		},
		{
			name:   "2",
			digits: []int{0, 2, 2},
			res:    2,
		},
		{
			name:   "3",
			digits: []int{6, 6, 6},
			res:    1,
		},
		{
			name:   "4",
			digits: []int{1, 3, 5},
			res:    0,
		},
		{
			name:   "5",
			digits: []int{1, 2, 3},
			res:    2,
		},
		{
			name:   "6",
			digits: []int{1, 2, 4},
			res:    4,
		},
		{
			name:   "7",
			digits: []int{1, 2, 0, 4},
			res:    14,
		},
		{
			name:   "7_1",
			digits: []int{1, 2, 6, 4},
			res:    18,
		},
		{
			name:   "8",
			digits: []int{1, 3, 0, 4},
			res:    10,
		},
		{
			name:   "9",
			digits: []int{0, 1, 2},
			res:    3,
		},
		{
			name:   "10",
			digits: []int{4, 1, 2},
			res:    4,
		},
		{
			name:   "11",
			digits: []int{0, 1, 3},
			res:    2,
		},
		{
			name:   "12",
			digits: []int{0, 0, 1, 3},
			res:    4,
		},
		{
			name:   "13",
			digits: []int{1, 2, 4, 6, 8},
			res:    48,
		},
		{
			name:   "14",
			digits: []int{1, 2, 4, 6, 0},
			res:    39,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, totalNumbers(c.digits))
		})
	}
}

func totalNumbers(digits []int) int {
	result := make(map[int]struct{})
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			for k := 0; k < len(digits); k++ {
				if digits[i] == 0 || digits[k]%2 != 0 {
					continue
				}
				if i == k || j == k || i == j {
					continue
				}

				key := digits[i]*100 + digits[j]*10 + digits[k]
				result[key] = struct{}{}
			}
		}
	}

	return len(result)
}
