package _338_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestCountingBits(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  []int
	}{
		{
			name: "1",
			n:    2,
			res:  []int{0, 1, 1},
		},
		{
			name: "2",
			n:    5,
			res:  []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "3",
			n:    0,
			res:  []int{0},
		},
		{
			name: "4",
			n:    1,
			res:  []int{0, 1},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := countBits(c.n)
			assert.Equal(t, c.res, res)
		})
	}
}

func countBits(n int) []int {
	out := make([]int, 0, n+1)
	const sub string = "1"

	for i := 0; i <= n; i++ {
		s := strconv.FormatInt(int64(i), 2)
		c := strings.Count(s, sub)
		out = append(out, c)
	}
	return out
}
