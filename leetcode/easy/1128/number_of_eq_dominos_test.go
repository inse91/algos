package _1128_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/number-of-equivalent-domino-pairs/description/

func Test(t *testing.T) {
	cases := []struct {
		name     string
		dominoes [][]int
		res      int
	}{
		{
			name:     "1",
			dominoes: [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}},
			res:      1,
		},
		{
			name:     "2",
			dominoes: [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}},
			res:      3,
		},
		{
			name:     "3",
			dominoes: [][]int{{2, 1}, {1, 2}, {1, 2}, {1, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}},
			res:      15,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numEquivDominoPairs(c.dominoes))
		})
	}
}

func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[[2]int]int, len(dominoes))

	for _, d := range dominoes {
		_, ok := m[[2]int{d[0], d[1]}]
		if ok {
			m[[2]int{d[0], d[1]}]++
			continue
		}

		_, ok = m[[2]int{d[1], d[0]}]
		if ok {
			m[[2]int{d[1], d[0]}]++
			continue
		}

		m[[2]int{d[0], d[1]}] = 1
	}

	var res int
	for _, q := range m {
		if q <= 1 {
			continue
		}

		res += fact(q)
	}

	return res
}

func fact(n int) int {
	var res int
	for i := 1; i < n; i++ {
		res += i
	}

	return res
}
