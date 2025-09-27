package _2951_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-words-containing-character

func Test(t *testing.T) {
	cases := []struct {
		name string
		mntn []int
		res  []int
	}{
		{
			name: "1",
			mntn: []int{2, 4, 4},
			res:  []int{},
		},
		{
			name: "2",
			mntn: []int{1, 4, 3, 8, 5},
			res:  []int{1, 3},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, findPeaks(c.mntn))
		})
	}
}

func findPeaks(mountain []int) []int {
	var peaks []int
	for i := 1; i < len(mountain)-1; i++ {
		prev, cur := mountain[i-1], mountain[i]
		if cur <= prev {
			continue
		}

		next := mountain[i+1]
		if cur > next {
			peaks = append(peaks, i)
		}
	}

	return peaks
}
