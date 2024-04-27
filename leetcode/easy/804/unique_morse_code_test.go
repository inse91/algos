package _804_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUniqueMorseCode(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   int
	}{
		{
			name:  "1",
			words: []string{"gin", "zen", "gig", "msg"},
			res:   2,
		},
		{
			name:  "2",
			words: []string{"a"},
			res:   1,
		},
	}

	a := 'a'
	_ = a
	b := 'b'
	_ = b
	z := 'z'
	_ = z
	ma := morse[a-97]
	mz := morse[z-97]
	_ = ma
	_ = mz

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, uniqueMorseRepresentations(c.words))
		})
	}
}

var morse = [26]string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]struct{}, len(words))
	for _, w := range words {
		_ = w
		_ = m
		morseCode := wordToMorse(w)
		m[morseCode] = struct{}{}
	}

	return len(m)
}

func wordToMorse(w string) string {
	const gap = 97
	var sb strings.Builder
	for _, r := range w {
		sb.WriteString(morse[r-gap])
	}

	return sb.String()
}
