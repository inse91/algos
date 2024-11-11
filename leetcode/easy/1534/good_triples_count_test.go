package _1534_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-good-triplets/description/

func Test(t *testing.T) {
	cases := []struct {
		name    string
		arr     []int
		a, b, c int
		res     int
	}{
		{
			name: "1",
			arr:  []int{3, 0, 1, 1, 9, 7},
			a:    7,
			b:    2,
			c:    3,
			res:  4,
		},
		{
			name: "2",
			arr:  []int{1, 1, 2, 2, 3},
			a:    0,
			b:    0,
			c:    1,
			res:  0,
		},
		{
			name: "12",
			arr:  []int{7, 3, 7, 3, 12, 1, 12, 2, 3},
			a:    5,
			b:    8,
			c:    1,
			res:  12,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countGoodTriplets(c.arr, c.a, c.b, c.c))
		})
	}
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				_i, _j, _k := arr[i], arr[j], arr[k]
				_a := mod(_i, _j)
				_b := mod(_j, _k)
				_c := mod(_k, _i)

				if _a <= a && _b <= b && _c <= c {
					res++
				}
			}
		}
	}

	return res
}

func mod(a, b int) int {
	d := a - b
	if d >= 0 {
		return d
	}

	return -d
}
