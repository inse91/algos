package _500_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestKeyBoardRow(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   []string
	}{
		{
			name:  "1",
			words: []string{"Hello", "Alaska", "Dad", "Peace"},
			res:   []string{"Alaska", "Dad"},
		},
		{
			name:  "2",
			words: []string{"omk"},
			res:   []string{},
		},
		{
			name:  "3",
			words: []string{"adsdf", "sfd"},
			res:   []string{"adsdf", "sfd"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findWords(c.words)
			assert.Equal(t, c.res, res)
		})
	}
}

const firstRow = "qwertyuiop"
const secondRow = "asdfghjkl"
const thirdRow = "zxcvbnm"

func findWords(words []string) []string {
	res := make([]string, 0, len(words))
	for _, w := range words {
		if !checkWord(strings.ToLower(w)) {
			continue
		}

		res = append(res, w)
	}
	return res
}

func checkWord(word string) bool {
	if len(word) == 0 {
		return true
	}
	firstLetter := string(word[0])
	switch {
	case strings.Contains(firstRow, firstLetter):
		return check(word, firstRow)
	case strings.Contains(secondRow, firstLetter):
		return check(word, secondRow)
	case strings.Contains(thirdRow, firstLetter):
		return check(word, thirdRow)
	}
	return true
}

func check(s string, letters string) bool {
	for _, r := range s {
		if !strings.Contains(letters, string(r)) {
			return false
		}
	}
	return true
}
