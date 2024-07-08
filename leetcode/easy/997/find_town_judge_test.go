package _997_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/find-the-town-judge/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		n     int
		trust [][]int
		res   int
	}{
		{
			name:  "1",
			n:     2,
			trust: [][]int{{1, 2}},
			res:   2,
		},
		{
			name:  "2",
			n:     3,
			trust: [][]int{{1, 3}, {2, 3}},
			res:   3,
		},
		{
			name:  "3",
			n:     3,
			trust: [][]int{{1, 3}, {2, 3}, {3, 1}},
			res:   -1,
		},
		{
			name:  "90",
			n:     1,
			trust: [][]int{},
			res:   1,
		},
		{
			name:  "91",
			n:     2,
			trust: [][]int{},
			res:   -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findJudge(c.n, c.trust))
		})
	}
}

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 {
		if n == 1 {
			return 1
		}
		return -1
	}

	peopleSet := make(map[int]struct{}, n)
	trustMap := make(map[int]int, n)

	for _, tr := range trust {
		whoTrust := tr[0]
		whomTrust := tr[1]
		peopleSet[whoTrust] = struct{}{}
		trustMap[whomTrust]++
	}

	var pj int
	for i := 1; i <= n; i++ {
		_, ok := peopleSet[i]
		if !ok {
			pj = i
		}
	}
	if pj == 0 {
		return -1
	}

	v, ok := trustMap[pj]
	if !ok {
		return -1
	}

	if v != n-1 {
		return -1
	}

	return pj
}
