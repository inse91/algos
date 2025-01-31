package _1768_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-strings-alternately/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   string
	}{
		{
			name:  "1",
			words: []string{"abc", "pqr"},
			res:   "apbqcr",
		},
		{
			name:  "2",
			words: []string{"ab", "pqrs"},
			res:   "apbqrs",
		},
		{
			name:  "3",
			words: []string{"abcd", "pq"},
			res:   "apbqcd",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, mergeAlternately(c.words[0], c.words[1]))
		})
	}
}

func mergeAlternately(word1 string, word2 string) string {
	var (
		res    strings.Builder
		c1, c2 bool
	)

	for i, j := 0, 0; ; {
		if !c1 {
			c1 = i == len(word1)
		}
		if !c2 {
			c2 = j == len(word2)
		}

		if c1 && c2 {
			break
		}

		if !c1 {
			res.WriteByte(word1[i])
			i++
		}
		if !c2 {
			res.WriteByte(word2[j])
			j++
		}
	}

	return res.String()
}
