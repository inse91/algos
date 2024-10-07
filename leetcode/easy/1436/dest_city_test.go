package _1436_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lucky-numbers-in-a-matrix/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		paths [][]string
		res   string
	}{
		{
			name:  "1",
			paths: [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			res:   "Sao Paulo",
		},
		{
			name:  "2",
			paths: [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			res:   "A",
		},
		{
			name:  "3",
			paths: [][]string{{"A", "Z"}},
			res:   "Z",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, destCity(c.paths))
		})
	}
}

func destCity(paths [][]string) string {
	deps := make(map[string]struct{}, len(paths))
	dests := make(map[string]struct{}, len(paths))

	for _, p := range paths {
		if len(p) < 2 {
			continue
		}
		deps[p[0]] = struct{}{}
		dests[p[1]] = struct{}{}
	}

	for d := range dests {
		_, ok := deps[d]
		if !ok {
			return d
		}
	}

	return ""
}
