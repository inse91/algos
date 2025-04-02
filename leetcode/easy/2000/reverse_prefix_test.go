package _2000_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-prefix-of-word/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		word string
		ch   byte
		res  string
	}{
		{
			name: "1",
			word: "abcdefd",
			ch:   'd',
			res:  "dcbaefd",
		},
		{
			name: "2",
			word: "xyxzxe",
			ch:   'z',
			res:  "zxyxxe",
		},
		{
			name: "3",
			word: "abcd",
			ch:   'z',
			res:  "abcd",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, reversePrefix(c.word, c.ch))
		})
	}
}

func reversePrefix(word string, ch byte) string {
	var sb strings.Builder
	for i, r := range word {
		b := byte(r)
		if b != ch {
			continue
		}

		sb.Grow(len(word))
		start := []byte(word[:i+1])
		slices.Reverse(start)
		_, err := sb.WriteString(string(start))
		if err != nil {
			panic(err)
		}

		sb.WriteString(word[i+1:])

		return sb.String()
	}

	return word
}
