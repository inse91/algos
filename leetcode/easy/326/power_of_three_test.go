package _326_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowerOfThree(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  bool
	}{
		{
			name: "1",
			in:   27,
			res:  true,
		},
		{
			name: "2",
			in:   0,
			res:  false,
		},
		{
			name: "3",
			in:   9,
			res:  true,
		},
		{
			name: "4",
			in:   45,
			res:  false,
		},
		{
			name: "5",
			in:   1,
			res:  true,
		},
		{
			name: "6",
			in:   -1,
			res:  false,
		},
		{
			name: "7",
			in:   3,
			res:  true,
		},
		{
			name: "8",
			in:   81,
			res:  true,
		},
		{
			name: "9",
			in:   82,
			res:  false,
		},
		{
			name: "10",
			in:   4,
			res:  false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := isPowerOfThree(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n >= 3 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return n == 1
}

func sumOfDigits(n int) (sum int) {
	const q = 10
	for n > 0 {
		sum += n - (n/q)*q
		n /= q
	}
	return sum
}

func TestSumOfDigits(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  int
	}{
		{
			name: "1",
			in:   27,
			res:  9,
		},
		{
			name: "2",
			in:   13,
			res:  4,
		},
		{
			name: "3",
			in:   111,
			res:  3,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := sumOfDigits(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}
