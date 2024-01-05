package _345_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	cases := []struct {
		name string
		in   string
		res  string
	}{
		{
			name: "1",
			in:   "hello",
			res:  "holle",
		},
		{
			name: "2",
			in:   "leetcode",
			res:  "leotcede",
		},
		{
			name: "3",
			in:   "aA",
			res:  "Aa",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := reverseVowels(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

var vowels = map[rune]struct{}{
	'a': {},
	'A': {},
	'e': {},
	'E': {},
	'i': {},
	'I': {},
	'o': {},
	'O': {},
	'u': {},
	'U': {},
}

func isVowel(r rune) (ok bool) {
	_, ok = vowels[r]
	return ok
}

func reverseVowels(s string) string {
	//vowelPlaces := make(map[int]rune)
	vowelPlaces := make([]byte, 0)

	for _, r := range s {
		if !isVowel(r) {
			continue
		}
		vowelPlaces = append(vowelPlaces, byte(r))
	}

	bts := []byte(s)

	c := len(vowelPlaces) - 1

	for i, b := range bts {
		if !isVowel(rune(b)) {
			continue
		}
		bts[i] = vowelPlaces[c]
		c--
	}

	return string(bts)
}
