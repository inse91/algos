package _18

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	cases := []struct {
		name string
		val  int
		res  [][]int
	}{
		{
			name: "1",
			val:  1,
			res:  [][]int{{1}},
		},
		{
			name: "2",
			val:  2,
			res:  [][]int{{1}, {1, 1}},
		},
		{
			name: "3",
			val:  3,
			res:  [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			name: "4",
			val:  4,
			res:  [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, generate(c.val))
		})
	}
}

func generate(numRows int) [][]int {
	out := make([][]int, 0, numRows)
	v := []int{1}
	out = append(out, v)
	for i := 2; i <= numRows; i++ {
		sl := makeSlice(i, v)
		v = sl
		out = append(out, sl)
	}

	return out
}

func makeSlice(n int, prev []int) []int {
	out := make([]int, n)
	out[0] = 1
	out[len(out)-1] = 1
	for i := 1; i < n-1; i++ {
		out[i] = prev[i-1] + prev[i]
	}
	return out
}
