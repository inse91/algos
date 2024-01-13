package _412_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  []string
	}{
		{
			name: "1",
			in:   3,
			res:  []string{"1", "2", "Fizz"},
		},
		{
			name: "2",
			in:   5,
			res:  []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "3",
			in:   15,
			res:  []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := fizzBuzz(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func fizzBuzz(n int) []string {

	res := make([]string, 0, n)
	const div5 = "Buzz"
	const div3 = "Fizz"
	const div3and5 = "FizzBuzz"

	var s string

	for i := 1; i <= n; i++ {
		divBy5 := i%5 == 0
		divBy3 := i%3 == 0
		switch {
		case divBy5 && divBy3:
			s = div3and5
		case divBy5:
			s = div5
		case divBy3:
			s = div3
		default:
			s = strconv.Itoa(i)
		}
		res = append(res, s)
	}

	return res
}
