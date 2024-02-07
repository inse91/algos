package _506_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"strconv"
	"testing"
)

func TestRelativeRanks(t *testing.T) {
	cases := []struct {
		name  string
		score []int
		res   []string
	}{
		{
			name:  "1",
			score: []int{5, 4, 3, 2, 1},
			res:   []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			name:  "2",
			score: []int{10, 3, 8, 9, 4},
			res:   []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
		{
			name:  "3",
			score: []int{10},
			res:   []string{"Gold Medal"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findRelativeRanks(c.score)
			assert.Equal(t, c.res, res)
		})
	}
}

func findRelativeRanks(score []int) []string {
	scoreCopy := make([]int, len(score))
	copy(scoreCopy, score)
	m := make(map[int]string, len(score))

	const first = "Gold Medal"
	const sec = "Silver Medal"
	const third = "Bronze Medal"

	slices.SortFunc(scoreCopy, func(a int, b int) int {
		if a > b {
			return -1
		}
		return 1
	})

	for i, num := range scoreCopy[:min(len(score), 3)] {
		switch i {
		case 0:
			m[num] = first
		case 1:
			m[num] = sec
		case 2:
			m[num] = third
		default:
		}
	}

	if len(score) >= 4 {
		for i, num := range scoreCopy[3:] {
			m[num] = strconv.Itoa(i + 4)
		}
	}

	res := make([]string, 0, len(score))

	for _, num := range score {
		res = append(res, m[num])
	}

	return res
}
