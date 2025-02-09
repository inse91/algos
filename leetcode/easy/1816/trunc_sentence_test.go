package _1816_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://leetcode.com/problems/truncate-sentence/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		k    int
		res  string
	}{
		{
			name: "1",
			s:    "Hello how are you Contestant",
			k:    4,
			res:  "Hello how are you",
		},
		{
			name: "2",
			s:    "What is the solution to this problem",
			k:    4,
			res:  "What is the solution",
		},
		{
			name: "3",
			s:    "chopper is not a tanuki",
			k:    5,
			res:  "chopper is not a tanuki",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, truncateSentence(c.s, c.k))
		})
	}
}

func truncateSentence(s string, k int) string {
	slc := strings.Split(s, " ")
	if len(slc) == k {
		return s
	}

	return strings.Join(slc[:k], " ")
}
