package _832_test

// https://leetcode.com/problems/flipping-an-image/description/

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		img  [][]int
		res  [][]int
	}{
		{
			name: "1",
			img:  [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			res:  [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			name: "2",
			img:  [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			res:  [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := flipAndInvertImage(c.img)
			for i, r := range res {
				assert.ElementsMatch(t, c.res[i], r)
			}
		})
	}
}

func flipAndInvertImage(image [][]int) [][]int {
	for i := range image {
		flipRow(image[i])
	}

	return image
}

func flipRow(row []int) {
	var l, r int
	for i, j := 0, len(row)-1; i <= j; func() { i++; j-- }() {
		if i == j {
			if row[i] == 0 {
				row[i] = 1
			} else {
				row[i] = 0
			}
			break
		}
		l, r = row[i], row[j]
		if l == 0 {
			row[j] = 1
		} else {
			row[j] = 0
		}
		if r == 0 {
			row[i] = 1
		} else {
			row[i] = 0
		}
	}
}
