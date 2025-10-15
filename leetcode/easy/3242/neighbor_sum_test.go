package _3242_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/design-neighbor-sum-service

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		grid := [][]int{
			{0, 1, 2},
			{3, 4, 5},
			{6, 7, 8},
		}
		ns := Constructor(grid)

		assert.Equal(t, 6, ns.AdjacentSum(1))
		assert.Equal(t, 16, ns.AdjacentSum(4))
		assert.Equal(t, 16, ns.DiagonalSum(4))
		assert.Equal(t, 4, ns.DiagonalSum(8))
	})

	t.Run("2", func(t *testing.T) {
		grid := [][]int{
			{1, 2, 0, 3},
			{4, 7, 15, 6},
			{8, 9, 10, 11},
			{12, 13, 14, 5},
		}
		ns := Constructor(grid)

		assert.Equal(t, 23, ns.AdjacentSum(15))
		assert.Equal(t, 45, ns.DiagonalSum(9))
	})

	t.Run("3", func(t *testing.T) {
		grid := [][]int{
			{3, 6, 0},
			{8, 5, 1},
			{2, 4, 7},
		}
		ns := Constructor(grid)

		assert.Equal(t, 12, ns.AdjacentSum(1))
		assert.Equal(t, 14, ns.AdjacentSum(3))
		assert.Equal(t, 5, ns.AdjacentSum(7))
		assert.Equal(t, 12, ns.AdjacentSum(2))
		assert.Equal(t, 19, ns.AdjacentSum(5))
		assert.Equal(t, 10, ns.AdjacentSum(8))
		assert.Equal(t, 7, ns.AdjacentSum(0))
		assert.Equal(t, 10, ns.DiagonalSum(1)) // bad
		assert.Equal(t, 14, ns.AdjacentSum(4))
		assert.Equal(t, 10, ns.DiagonalSum(8)) // bad
		assert.Equal(t, 8, ns.AdjacentSum(6))
		assert.Equal(t, 5, ns.DiagonalSum(3))
		assert.Equal(t, 5, ns.DiagonalSum(2))
		assert.Equal(t, 9, ns.DiagonalSum(6))
		assert.Equal(t, 12, ns.DiagonalSum(5))
		assert.Equal(t, 5, ns.DiagonalSum(7))
		assert.Equal(t, 9, ns.DiagonalSum(4))
		assert.Equal(t, 5, ns.DiagonalSum(0))
	})
}

type NeighborSum struct {
	grid [][]int
	pts  [100][2]int
	n    int
}

func Constructor(grid [][]int) NeighborSum {
	ns := NeighborSum{
		grid: grid,
		n:    len(grid) - 1,
	}
	for i := range grid {
		for j := range grid[i] {
			num := grid[i][j]
			ns.pts[num] = [2]int{i, j}
		}
	}

	return ns
}

func (this *NeighborSum) AdjacentSum(value int) int {
	var (
		top, right, down, left int
	)

	pt := this.pts[value]
	i, j := pt[0], pt[1]
	if i != 0 {
		top = this.grid[i-1][j]
	}
	if i != this.n {
		down = this.grid[i+1][j]
	}
	if j != 0 {
		left = this.grid[i][j-1]
	}
	if j != this.n {
		right = this.grid[i][j+1]
	}

	return top + right + down + left
}

func (this *NeighborSum) DiagonalSum(value int) int {
	pt := this.pts[value]
	i, j := pt[0], pt[1]
	if i == 0 {
		down := this.grid[i+1]
		if j == 0 {
			return down[j+1]
		}
		if j == this.n {
			return down[j-1]
		}

		return down[j-1] + down[j+1]
	}
	if i == this.n {
		top := this.grid[i-1]
		if j == 0 {
			return top[j+1]
		}
		if j == this.n {
			return top[j-1]
		}

		return top[j-1] + top[j+1]
	}
	if j == 0 {
		tr := this.grid[i-1][j+1]
		dr := this.grid[i+1][j+1]

		return tr + dr
	}
	if j == this.n {
		tl := this.grid[i-1][j-1]
		dl := this.grid[i+1][j-1]

		return tl + dl
	}

	return this.grid[i+1][j+1] + this.grid[i+1][j-1] + this.grid[i-1][j+1] + this.grid[i-1][j-1]
}
