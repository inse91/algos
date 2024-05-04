package _867_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		in   string
		res  string
	}{
		{
			name: "1",
			in:   "I speak Goat Latin",
			res:  "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			name: "2",
			in:   "The quick brown fox jumped over the lazy dog",
			res:  "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, toGoatLatin(c.in))
		})
	}
}

func toGoatLatin(sentence string) string {
	var sb strings.Builder
	sb.Grow(len(sentence) + 1)

	const ma = "ma"
	const a = 'a'
	const space = ' '
	q := 1

	words := strings.Split(sentence, " ")
	for _, w := range words {
		if len(w) == 0 {
			continue
		}

		if !isVowel(rune(w[0])) {
			w = w[1:] + w[:1]
		}

		sb.WriteString(w)
		sb.WriteString(ma)
		for i := 0; i < q; i++ {
			sb.WriteRune(a)
		}
		q++

		sb.WriteRune(space)
	}

	sentence = sb.String()
	sb.Reset()

	return sentence[:len(sentence)-1]
}

func isVowel(r rune) bool {
	_, ok := vowelSet[r]

	return ok
}

var vowelSet = map[rune]struct{}{
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
