package _2351_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/first-letter-to-appear-twice

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  byte
	}{
		{
			name: "1",
			s:    "abccbaacz",
			res:  'c',
		},
		{
			name: "2",
			s:    "abcdd",
			res:  'd',
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, repeatedCharacter(c.s))
		})
	}
}

func repeatedCharacter(s string) byte {
	var letters [26]bool
	for _, r := range s {
		idx := int(r - 'a')
		if letters[idx] {
			return byte(r)
		}

		letters[idx] = true
	}

	return 0
}
