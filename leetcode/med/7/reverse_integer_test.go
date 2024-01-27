package _7_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"slices"
	"strconv"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	cases := []struct {
		name string
		x    int
		res  int
	}{
		{
			name: "1",
			x:    123,
			res:  321,
		},
		{
			name: "2",
			x:    -123,
			res:  -321,
		},
		{
			name: "3",
			x:    120,
			res:  21,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := reverse(c.x)
			assert.Equal(t, c.res, res)
		})
	}
}

func reverse(x int) int {
	var isNegative bool
	if x < 0 {
		isNegative = true
		x = -x
	}

	s := strconv.Itoa(x)
	s = reverseStr(s)
	res, _ := strconv.Atoi(s)
	if res > math.MaxInt32 {
		return 0
	}
	if isNegative {
		return -res
	}
	return res
}

func reverseStr(s string) string {
	bts := []byte(s)
	slices.Reverse(bts)
	return string(bts)
}
