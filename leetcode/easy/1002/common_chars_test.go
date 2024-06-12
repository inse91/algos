package _1002_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://leetcode.com/problems/most-common-word/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   []string
	}{
		{
			name:  "1",
			words: []string{"bella", "label", "roller"},
			res:   []string{"e", "l", "l"},
		},
		{
			name:  "2",
			words: []string{"cool", "lock", "cook"},
			res:   []string{"c", "o"},
		},
		{
			name:  "3",
			words: []string{},
			res:   []string{},
		},
		{
			name:  "4",
			words: []string{"foo"},
			res:   []string{"f", "o", "o"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, commonChars(c.words))
		})
	}
}

func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	if len(words) == 1 {
		return strings.Split(words[0], "")
	}

	m := make(map[rune]int)
	for _, r := range words[0] {
		if _, ok := m[r]; !ok {
			m[r] = 1
			continue
		}
		m[r]++
	}

	for _, w := range words[1:] {
		wMap := make(map[rune]int)
		for _, r := range w {
			if _, ok := wMap[r]; !ok {
				wMap[r] = 1
				continue
			}
			wMap[r]++
		}

		var keysToDelete []rune
		for r, i := range m {
			iw, ok := wMap[r]
			if !ok {
				keysToDelete = append(keysToDelete, r)
				continue
			}

			m[r] = min(i, iw)
		}

		for _, r := range keysToDelete {
			delete(m, r)
		}
	}

	res := make([]string, 0, len(m))
	for r, i := range m {
		if i == 0 {
			continue
		}
		if i == 1 {
			res = append(res, string(r))
			continue
		}

		strs := strings.Split(
			strings.Repeat(string(r), i),
			"",
		)
		res = append(res, strs...)
	}

	return res
}
