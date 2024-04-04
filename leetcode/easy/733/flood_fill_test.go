package _733_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloodFill(t *testing.T) {
	cases := []struct {
		name  string
		image [][]int
		sr    int
		sc    int
		color int
		res   [][]int
	}{
		{
			name: "1",
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:    1,
			sc:    1,
			color: 2,
			res: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "2",
			image: [][]int{
				{0, 0, 0},
			},
			sr:    0,
			sc:    0,
			color: 0,
			res: [][]int{
				{0, 0, 0},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := floodFill(c.image, c.sr, c.sc, c.color)
			assert.Equal(t, c.res, res)
		})
	}
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	startColor := image[sr][sc]
	if startColor == color {
		return image
	}

	image[sr][sc] = color
	searchAndPaint(image, sr, sc, color, startColor)

	return image
}

func searchAndPaint(image [][]int, i, j, color, startColor int) {
	// down
	if i+1 < len(image) {
		func() {
			rv := image[i+1][j]
			if rv != startColor {
				return
			}
			image[i+1][j] = color
			searchAndPaint(image, i+1, j, color, startColor)
		}()
	}

	// up
	if i-1 >= 0 {
		func() {
			rv := image[i-1][j]
			if rv != startColor {
				return
			}
			image[i-1][j] = color
			searchAndPaint(image, i-1, j, color, startColor)
		}()
	}

	// left
	if j-1 >= 0 {
		func() {
			rv := image[i][j-1]
			if rv != startColor {
				return
			}
			image[i][j-1] = color
			searchAndPaint(image, i, j-1, color, startColor)
		}()
	}

	// right
	if j+1 < len(image[i]) {
		func() {
			rv := image[i][j+1]
			if rv != startColor {
				return
			}
			image[i][j+1] = color
			searchAndPaint(image, i, j+1, color, startColor)
		}()
	}
}
