package _830_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  [][]int
	}{
		{
			name: "1",
			s:    "abbxxxxzzy",
			res:  [][]int{{3, 6}},
		},
		{
			name: "1_1",
			s:    "abbxxxxzzzzz",
			res:  [][]int{{3, 6}, {7, 11}},
		},
		{
			name: "2",
			s:    "abc",
			res:  [][]int{},
		},
		{
			name: "3",
			s:    "abcdddeeeeaabbbcd",
			res:  [][]int{{3, 5}, {6, 9}, {12, 14}},
		},
		{
			name: "137",
			s:    "aaa",
			res:  [][]int{{0, 2}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := largeGroupPositions(c.s)
			for i, r := range res {
				assert.ElementsMatch(t, c.res[i], r)
			}
		})
	}
}

func largeGroupPositions(s string) [][]int {
	res := make([][]int, 0)

	c := 1
	for i := 1; i < len(s); i++ {
		si := string(s[i])
		sprev := string(s[i-1])
		_, _ = si, sprev
		if s[i] == s[i-1] {
			c++
			continue
		}
		if c >= 3 {
			res = append(res, []int{i - c, i - 1})
		}
		c = 1
	}
	if c >= 3 {
		i := len(s) - 1
		res = append(res, []int{i - c + 1, i})
	}

	return res
}
