package _2839_tes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-i

func Test(t *testing.T) {
	cases := []struct {
		name   string
		w1, w2 string
		res    bool
	}{
		{
			name: "1",
			w1:   "abcd",
			w2:   "cdab",
			res:  true,
		},
		{
			name: "2",
			w1:   "abcd",
			w2:   "dacb",
			res:  false,
		},
		{
			name: "607",
			w1:   "bnxw",
			w2:   "bwxn",
			res:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, canBeEqual(c.w1, c.w2))
		})
	}
}

func canBeEqual(s1 string, s2 string) bool {
	return ((s1[0] == s2[2] && s1[2] == s2[0]) || (s1[0] == s2[0] && s1[2] == s2[2])) &&
		(s1[1] == s2[3] && s1[3] == s2[1] || s1[1] == s2[1] && s1[3] == s2[3])
}
