package _1232_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		line [][]int
		res  bool
	}{
		{
			name: "1",
			line: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{5, 6},
				{6, 7},
			},
			res: true,
		},
		{
			name: "1_not_sorted",
			line: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{6, 7},
				{5, 6},
			},
			res: true,
		},
		{
			name: "2",
			line: [][]int{
				{1, 1},
				{2, 2},
				{3, 4},
				{4, 5},
				{5, 6},
				{7, 7},
			},
			res: false,
		},
		{
			name: "46",
			line: [][]int{
				{0, 0},
				{0, 1},
				{0, -1},
			},
			res: true,
		},
		{
			name: "48",
			line: [][]int{
				{2, 4},
				{2, 5},
				{2, 8},
			},
			res: true,
		},
		{
			name: "-1",
			line: [][]int{
				{0, 0},
				{1, 1},
				{2, 2},
			},
			res: true,
		},
		{
			name: "-2",
			line: [][]int{
				{0, 0},
				{1, 1},
				{2, 0},
			},
			res: false,
		},
		{
			name: "-3",
			line: [][]int{
				{1, 2},
				{2, 4},
				{4, 8},
			},
			res: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, checkStraightLine(c.line))
		})
	}
}

func checkStraightLine(line [][]int) bool {
	if len(line) < 2 {
		panic("wrong line")
	}

	dx := line[1][0] - line[0][0]
	dy := line[1][1] - line[0][1]

	vec := float64(dx) / float64(dy)

	if dx != 0 && dy != 0 {
		var ax, bx, ay, by int
		slices.SortFunc(line, func(a, b []int) int {
			ax, ay, bx, by = a[0], a[1], b[0], b[1]
			if dx == 0 {
				return ay - by
			}

			return ax - bx
		})
	}

	var point [2]int
	var ndX, ndY int
	prevPoint := [2]int{line[1][0], line[1][1]}
	for i := 2; i < len(line); i++ {
		point = [2]int{line[i][0], line[i][1]}
		ndX = point[0] - prevPoint[0]
		ndY = point[1] - prevPoint[1]

		if dx == 0 && ndX == 0 {
			continue
		}

		if dy == 0 && ndY == 0 {
			continue
		}

		if float64(ndX)/float64(ndY) != vec {
			return false
		}

		prevPoint = point
	}

	return true
}
