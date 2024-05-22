package _917_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "ab-cd",
			res:  "dc-ba",
		},
		{
			name: "2",
			s:    "a-bC-dEf-ghIj",
			res:  "j-Ih-gfE-dCba",
		},
		{
			name: "3",
			s:    "Test1ng-Leet=code-Q!",
			res:  "Qedo1ct-eeLg=ntse-T!",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, reverseOnlyLetters(c.s))
		})
	}
}

func reverseOnlyLetters(s string) string {
	spc := ' '
	letters := make([]rune, 0, len(s))

	var sb strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
			sb.WriteRune(spc)
			continue
		}
		sb.WriteRune(r)
	}

	s = sb.String()
	sb.Reset()
	idx := len(letters) - 1
	for _, r := range s {
		if r == spc {
			sb.WriteRune(letters[idx])
			idx--
			continue
		}
		sb.WriteRune(r)
	}

	return sb.String()
}
