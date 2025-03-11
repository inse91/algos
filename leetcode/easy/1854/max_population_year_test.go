package _1854_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-population-year/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		logs [][]int
		res  int
	}{
		{
			name: "1",
			logs: [][]int{{1993, 1999}, {2000, 2010}},
			res:  1993,
		},
		{
			name: "2",
			logs: [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}},
			res:  1960,
		},
		{
			name: "3",
			logs: [][]int{{1950, 1981}, {1962, 1963}, {1960, 1971}, {1970, 1981}},
			res:  1962,
		},
		{
			name: "35",
			logs: [][]int{{2008, 2026}, {2004, 2008}, {2034, 2035}, {1999, 2050}, {2049, 2050}, {2011, 2035}, {1966, 2033}, {2044, 2049}},
			res:  2011,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maximumPopulation(c.logs))
		})
	}
}

func maximumPopulation(logs [][]int) int {
	population := make(map[int]int)
	for _, log := range logs {
		if len(log) < 2 {
			panic("invalid log")
		}

		for yr := log[0]; yr < log[1]; yr++ {
			population[yr]++
		}
	}

	var year int
	var maxPopulation int
	for yr, pop := range population {
		if pop < maxPopulation {
			continue
		}

		if pop == maxPopulation {
			year = min(year, yr)
			continue
		}

		// pop > maxPopulation
		year = yr
		maxPopulation = pop
	}

	return year
}
