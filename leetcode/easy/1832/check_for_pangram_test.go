package _1832_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "thequickbrownfoxjumpsoverthelazydog",
			res:  true,
		},
		{
			name: "2",
			s:    "leetcode",
			res:  false,
		},
		{
			name: "3",
			s:    "az",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, checkIfPangram(c.s))
		})
	}
}

func checkIfPangram(sentence string) bool {
	const q = 26
	c := q
	var lettersCheck [q]bool // [false, false ... false]
	for _, r := range sentence {
		idx := r - 97
		if lettersCheck[idx] {
			continue
		}

		lettersCheck[idx] = true
		c--
	}

	return c == 0
}
