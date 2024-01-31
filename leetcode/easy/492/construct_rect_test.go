package _492_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestConstructRectangle(t *testing.T) {
	cases := []struct {
		name string
		area int
		res  []int
	}{
		{
			name: "1",
			area: 4,
			res:  []int{2, 2},
		},
		{
			name: "2",
			area: 37,
			res:  []int{37, 1},
		},
		{
			name: "3",
			area: 122122,
			res:  []int{427, 286},
		},
		{
			name: "4",
			area: 1,
			res:  []int{1, 1},
		},
		{
			name: "5",
			area: 16,
			res:  []int{4, 4},
		},
		{
			name: "6",
			area: 10,
			res:  []int{5, 2},
		},
		{
			name: "7",
			area: 12,
			res:  []int{4, 3},
		},
		{
			name: "8",
			area: 13,
			res:  []int{13, 1},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := constructRectangle(c.area)
			assert.Equal(t, c.res, res)
		})
	}

}

func constructRectangle(area int) []int {
	sq, ok := sqrt(area)
	if ok {
		return []int{sq, sq}
	}

	for l, w := sq+1, sq; w != 1; func() {
		l++
		w--
	}() {
		if l*w == area {
			return []int{l, w}
		}
		ll := l
		for {
			ll++
			if ll*w == area {
				return []int{ll, w}
			}
			if ll*w > area {
				break
			}
		}
	}

	return []int{area, 1}
}

func sqrt(v int) (int, bool) {
	sq_rt := int(math.Sqrt(float64(v)))
	return sq_rt, sq_rt*sq_rt == v
}
