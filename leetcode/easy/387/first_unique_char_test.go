package _387_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstUniqueChar(t *testing.T) {
	cases := []struct {
		name string
		in   string
		res  int
	}{
		{
			name: "1",
			in:   "leetcode",
			res:  0,
		},
		{
			name: "2",
			in:   "loveleetcode",
			res:  2,
		},
		{
			name: "3",
			in:   "aabb",
			res:  -1,
		},
		{
			name: "4",
			in:   "aa",
			res:  -1,
		},
		{
			name: "5",
			in:   "a",
			res:  0,
		},
		{
			name: "6",
			in:   "abaccdeff",
			res:  1,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := firstUniqChar(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func firstUniqChar(s string) int {

	chars := [26]struct {
		pos int
		qty int
	}{}
	//chars := make(map[byte]struct{}, 16)
	for i, r := range s {
		chars[r-'a'].pos = i
		chars[r-'a'].qty++
	}

	var got bool
	pos := len(s)
	for _, info := range chars {
		if info.qty == 1 {
			got = true
			if info.pos < pos {
				pos = info.pos
			}
		}
	}
	if !got {
		return -1
	}
	return pos
}
