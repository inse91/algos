package _520_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDetectCapital(t *testing.T) {
	cases := []struct {
		name string
		word string
		res  bool
	}{
		{
			name: "1",
			word: "USA",
			res:  true,
		},
		{
			name: "2",
			word: "FlaG",
			res:  false,
		},
		{
			name: "3",
			word: "Google",
			res:  true,
		},
		{
			name: "4",
			word: "google",
			res:  true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := detectCapitalUse(c.word)
			assert.Equal(t, c.res, res)
		})
	}
}

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	if strings.ToLower(word) == word || strings.ToUpper(word) == word {
		return true
	}

	rest := word[1:]
	return strings.ToLower(rest) == rest
}
