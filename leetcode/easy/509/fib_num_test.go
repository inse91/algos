package _509_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    2,
			res:  1,
		},
		{
			name: "2",
			n:    3,
			res:  2,
		},
		{
			name: "3",
			n:    4,
			res:  3,
		},
		{
			name: "4",
			n:    5,
			res:  5,
		},
		{
			name: "5",
			n:    6,
			res:  8,
		},
		{
			name: "6",
			n:    7,
			res:  13,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := fib(c.n)
			assert.Equal(t, c.res, res)
		})
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
