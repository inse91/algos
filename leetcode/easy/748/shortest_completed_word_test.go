package _748_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode"
)

func TestShortestCompletedWord(t *testing.T) {
	cases := []struct {
		name  string
		lic   string
		words []string
		res   string
	}{
		{
			name:  "1",
			lic:   "1s3 PSt",
			words: []string{"step", "steps", "stripe", "stepple"},
			res:   "steps",
		},
		{
			name:  "2",
			lic:   "1s3 456",
			words: []string{"looks", "pest", "stew", "show"},
			res:   "pest",
		},
		{
			name:  "86",
			lic:   "GrC8950",
			words: []string{"according", "measure", "other", "every", "base", "level", "meeting", "none", "marriage", "rest"},
			res:   "according",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, shortestCompletingWord(c.lic, c.words))
		})
	}
}

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)

	m := make(map[rune]int)
	for _, r := range licensePlate {
		s := string(r)
		_ = s
		if !unicode.IsLetter(r) {
			continue
		}

		if _, ok := m[r]; ok {
			m[r] += 1
			continue
		}

		m[r] = 1
	}

	var res string

	for _, word := range words {
		ok := true
		for r, q := range m {
			c := strings.Count(word, string(r))
			if c >= q {
				continue
			}
			ok = false
			break
		}

		if !ok {
			continue
		}
		if res == "" {
			res = word
		}

		res = min(res, word)

		if len(word) < len(res) {
			res = word
		}
	}

	return res
}
