package _2558_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/split-with-minimum-sum

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  4325,
			res:  59,
		},
		{
			name: "2",
			num:  687,
			res:  75,
		},
		{
			name: "11",
			num:  444325,
			res:  589,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, splitNum(c.num))
		})
	}
}

func splitNum(num int) int {
	m := make(map[int]int)
	for num > 0 {
		dig := num - num/10*10
		m[dig]++

		num /= 10
	}

	var n1, n2 int
	var state bool
	for i := range 10 {
		q := m[i]
		if q == 0 {
			continue
		}

		for range q {
			state = !state
			if state {
				n1 = n1*10 + i
				continue
			}

			n2 = n2*10 + i
		}
	}

	return n1 + n2
}
