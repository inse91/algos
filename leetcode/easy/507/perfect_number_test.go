package _507_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPerfectNumber(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  bool
	}{
		{
			name: "1",
			num:  28,
			res:  true,
		},
		{
			name: "2",
			num:  6,
			res:  true,
		},
		{
			name: "3",
			num:  7,
			res:  false,
		},
		{
			name: "4",
			num:  1,
			res:  false,
		},
		{
			name: "5",
			num:  2,
			res:  false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := checkPerfectNumber(c.num)
			assert.Equal(t, c.res, res)
		})
	}
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1

	limit := num / 2
	limit = int(math.Sqrt(float64(num)))

	for divisor := 2; divisor <= limit; divisor++ {
		if num%divisor == 0 {
			sum += divisor
			if num/divisor != divisor {
				sum += num / divisor
			}
		}
	}

	return sum == num
}
