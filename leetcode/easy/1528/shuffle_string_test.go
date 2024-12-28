package _1528_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/shuffle-string/description/

func Test(t *testing.T) {
	cases := []struct {
		name    string
		s       string
		indices []int
		res     string
	}{
		{
			name:    "1",
			s:       "codeleet",
			indices: []int{4, 5, 6, 7, 0, 2, 1, 3},
			res:     "leetcode",
		},
		{
			name:    "2",
			s:       "abc",
			indices: []int{0, 1, 2},
			res:     "abc",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, restoreString(c.s, c.indices))
		})
	}
}

func restoreString(s string, indices []int) string {
	bts := make([]byte, len(s))
	for i := range bts {
		bts[indices[i]] = s[i]
	}

	return string(bts)
}
