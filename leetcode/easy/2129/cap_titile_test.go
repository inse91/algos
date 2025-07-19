package _2129_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/capitalize-the-title

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "capitaLize thE tiTle",
			res:  "Capitalize The Title",
		},
		{
			name: "2",
			s:    "First leTTer of EACH Word",
			res:  "First Letter of Each Word",
		},
		{
			name: "3",
			s:    "i lOve leetcode",
			res:  "i Love Leetcode",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, capitalizeTitle(c.s))
		})
	}
}

func capitalizeTitle(title string) string {
	parts := strings.Split(title, " ")
	for i, pt := range parts {
		parts[i] = strings.ToLower(pt)
		if len(pt) <= 2 {
			continue
		}

		firstLetter := unicode.ToUpper(rune(pt[0]))
		parts[i] = string(firstLetter) + parts[i][1:]
	}

	return strings.Join(parts, " ")
}
