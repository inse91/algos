package _1935_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-number-of-words-you-can-type

func Test(t *testing.T) {
	cases := []struct {
		name   string
		text   string
		broken string
		res    int
	}{
		{
			name:   "1",
			text:   "hello world",
			broken: "ad",
			res:    1,
		},
		{
			name:   "2",
			text:   "leet code",
			broken: "lt",
			res:    1,
		},
		{
			name:   "3",
			text:   "leet code",
			broken: "e",
			res:    0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, canBeTypedWords(c.text, c.broken))
		})
	}
}

func canBeTypedWords(text string, brokenLetters string) int {
	m := make(map[rune]struct{}, len(brokenLetters))
	for _, r := range brokenLetters {
		m[r] = struct{}{}
	}

	var res int
	for _, w := range strings.Split(text, " ") {
		var invalid bool
		for _, r := range w {
			if _, ok := m[r]; !ok {
				continue
			}

			invalid = true
			break
		}

		if invalid {
			continue
		}

		res++
	}

	return res
}
