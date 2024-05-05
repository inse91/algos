package _24_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranspose(t *testing.T) {
	testCases := []struct {
		name string
		in   [][]int
		out  [][]int
	}{
		{
			name: "1",
			in:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			out:  [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			name: "2",
			in:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			out:  [][]int{{1, 5, 9}, {2, 6, 10}, {3, 7, 11}, {4, 8, 12}},
		},
		{
			name: "3",
			in:   [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}},
			out:  [][]int{{1, 6, 11}, {2, 7, 12}, {3, 8, 13}, {4, 9, 14}, {5, 10, 15}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.out, transpose(tc.in))
		})
	}
}

func transpose(matrix [][]int) [][]int {
	out := make([][]int, len(matrix[0]))
	for i := range matrix {
		for j, num := range matrix[i] {
			if out[j] == nil {
				out[j] = make([]int, 0, len(matrix))
			}
			out[j] = append(out[j], num)
		}
	}
	return out
}
