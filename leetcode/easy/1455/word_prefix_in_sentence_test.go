package _1455_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		stc  string
		wrd  string
		res  int
	}{
		{
			name: "1",
			stc:  "i love eating burger",
			wrd:  "burg",
			res:  4,
		},
		{
			name: "2",
			stc:  "this problem is an easy problem",
			wrd:  "pro",
			res:  2,
		},
		{
			name: "3",
			stc:  "i am tired",
			wrd:  "you",
			res:  -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isPrefixOfWord(c.stc, c.wrd))
		})
	}
}

func isPrefixOfWord(sentence string, searchWord string) int {
	const space = " "
	words := strings.Split(sentence, space)

	for i, word := range words {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}

	return -1
}
