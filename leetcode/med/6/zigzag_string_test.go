package _6_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestZigzagString(t *testing.T) {

	testCases := []struct {
		name string
		in   string
		rows int
		res  string
	}{
		{
			name: "1",
			in:   "PAYPALISHIRING",
			rows: 3,
			res:  "PAHNAPLSIIGYIR",
			// P   A   H   N
			// A P L S I I G
			// Y   I   R
		},
		{
			name: "2",
			in:   "PAYPALISHIRING",
			rows: 4,
			res:  "PINALSIGYAHRPI",
			// P     I    N
			// A   L S  I G
			// Y A   H R
			// P     I
		},
		{
			name: "3",
			in:   "PAYPALISHIRING",
			rows: 5,
			res:  "PHASIYIRPLIGAN",
			// P     H
			// A   S I
			// Y  I  R
			// P L   I G
			// A     N
		},
		{
			name: "4",
			rows: 1,
			in:   "A",
			res:  "A",
			// A
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := convert(tc.in, tc.rows)
			assert.Equal(t, tc.res, res)
		})
	}
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows == len(s) {
		return s
	}
	lenS := len(s)
	seq, maxIdx := getSeq(numRows)
	numOfIters := lenS / maxIdx

	sb := strings.Builder{}
	sb.Grow(lenS)

	for _, ints := range seq {
		for i := 0; i <= numOfIters; i++ {
			for _, idx := range ints {
				curIdx := maxIdx*i + idx
				if curIdx >= lenS {
					continue
				}
				sb.WriteByte(s[curIdx])
				fmt.Println(sb.String())
			}
		}
	}

	return sb.String()
}

func getSeq(numRows int) ([][]int, int) {
	seq := make([][]int, numRows)
	l := numRows - 1
	maxIdx := l * 2

	seq[0] = []int{0}
	seq[l] = []int{l}

	for i := 1; i < l; i++ {
		seq[i] = []int{i, maxIdx - i}
	}
	return seq, maxIdx
}
