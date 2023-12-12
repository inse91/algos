package _258_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestAddDigits(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		res  int
	}{
		{
			name: "1",
			in:   38,
			res:  2,
		},
		{
			name: "2",
			in:   0,
			res:  0,
		},
		{
			name: "3",
			in:   9,
			res:  9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := addDigits(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	nums := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	for num >= 10 {
		s := strconv.Itoa(num)
		num = 0
		for _, r := range s {
			num += nums[byte(r)]
		}
	}

	return num
}
