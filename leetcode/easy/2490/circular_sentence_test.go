package _2490_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/circular-sentence

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "leetcode exercises sound delightful",
			res:  true,
		},
		{
			name: "2",
			s:    "eetcode",
			res:  true,
		},
		{
			name: "3",
			s:    "Leetcode is cool",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isCircularSentence(c.s))
		})
	}
}

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if i == len(words)-1 {
			continue
		}

		next := words[i+1]
		if word[len(word)-1] != next[0] {
			return false
		}
	}

	// last word last letter compared to first word first letter
	last := words[len(words)-1]
	return words[0][0] == last[len(last)-1]
}
