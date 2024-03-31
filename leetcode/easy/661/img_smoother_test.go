package _661_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImgSmoother(t *testing.T) {
	cases := []struct {
		name string
		img  [][]int
		res  [][]int
	}{
		{
			name: "1",
			img: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			res: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "2",
			img: [][]int{
				{100, 200, 100},
				{200, 50, 200},
				{100, 200, 100},
			},
			res: [][]int{
				{137, 141, 137},
				{141, 138, 141},
				{137, 141, 137},
			},
		},
		{
			name: "3",
			img: [][]int{
				{1},
			},
			res: [][]int{
				{1},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := imageSmoother(c.img)
			assert.Equal(t, c.res, res)
		})
	}
}

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	n := len(img[0])

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i := range res {
		for j := range res[i] {
			c := 1

			// self
			sum := img[i][j]

			// top left
			if i-1 >= 0 && j-1 >= 0 {
				c++
				sum += img[i-1][j-1]
			}

			// top
			if i-1 >= 0 {
				c++
				sum += img[i-1][j]
			}

			// top right
			if i-1 >= 0 && j+1 < len(img[i]) {
				c++
				sum += img[i-1][j+1]
			}

			// left
			if j-1 >= 0 {
				c++
				sum += img[i][j-1]
			}

			// right
			if j+1 < len(img[i]) {
				c++
				sum += img[i][j+1]
			}

			// down left
			if i+1 < len(img) && j-1 >= 0 {
				c++
				sum += img[i+1][j-1]
			}

			// down
			if i+1 < len(img) {
				c++
				sum += img[i+1][j]
			}

			// down right
			if i+1 < len(img) && j+1 < len(img[i]) {
				c++
				sum += img[i+1][j+1]
			}

			res[i][j] = sum / c
		}
	}

	return res
}
