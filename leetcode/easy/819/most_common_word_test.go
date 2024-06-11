package _819_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/most-common-word/description/

func Test(t *testing.T) {
	cases := []struct {
		name      string
		paragraph string
		banned    []string
		res       string
	}{
		{
			name:      "1",
			paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned:    []string{"hit"},
			res:       "ball",
		},
		{
			name:      "2",
			paragraph: "a.",
			banned:    []string{},
			res:       "a",
		},
		{
			name:      "46",
			paragraph: "Bob",
			banned:    []string{},
			res:       "bob",
		},
		{
			name:      "47",
			paragraph: "a, a, a, a, b,b,b,c, c",
			banned:    []string{"a"},
			res:       "b",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, mostCommonWord(c.paragraph, c.banned))
		})
	}
}

func split(p string) []string {
	res := make([]string, 0, len(p)/4)

	var sb strings.Builder
	for _, r := range p {
		if unicode.IsLetter(r) {
			sb.WriteRune(unicode.ToLower(r))
			continue
		}

		if sb.Len() != 0 {
			res = append(res, sb.String())
			sb.Reset()
		}
	}

	if sb.Len() != 0 {
		res = append(res, sb.String())
	}

	return res
}

func mostCommonWord(paragraph string, banned []string) string {
	banSet := make(map[string]struct{}, len(banned))
	for _, w := range banned {
		banSet[w] = struct{}{}
	}

	words := split(paragraph)
	wordsCounts := make(map[string]int, len(words)/2)

	for _, w := range words {
		if _, ok := banSet[w]; ok {
			continue
		}

		c, ok := wordsCounts[w]
		if !ok {
			wordsCounts[w] = 1
			continue
		}
		wordsCounts[w] = c + 1
	}

	mx := -1
	var res string
	for w, c := range wordsCounts {
		if c <= mx {
			continue
		}
		mx = c
		res = w
	}

	return res
}
