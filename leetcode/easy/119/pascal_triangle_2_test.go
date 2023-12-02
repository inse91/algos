package _18

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPascalTriangle2(t *testing.T) {
	cases := []struct {
		name string
		val  int
		res  []int
	}{
		{
			name: "0",
			val:  0,
			res:  []int{1},
		},
		{
			name: "1",
			val:  1,
			res:  []int{1, 1},
		},
		{
			name: "2",
			val:  2,
			res:  []int{1, 2, 1},
		},
		{
			name: "3",
			val:  3,
			res:  []int{1, 3, 3, 1},
		},
		{
			name: "4",
			val:  4,
			res:  []int{1, 4, 6, 4, 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, getRow(c.val))
		})
	}
}

func getRow(numRows int) []int {
	out := make([]int, 0, numRows+1)

	out = append(out, 1)
	if numRows == 0 {
		return out
	}

	out = append(out, 1)
	if numRows == 1 {
		return out
	}
	var t int
	for i := 2; i <= numRows; i++ {
		var temp int = out[0]
		for j := 1; j < i; j++ {
			t = out[j]
			out[j] = out[j] + temp
			temp = t
		}
		out = append(out, 1)
	}

	return out
}
