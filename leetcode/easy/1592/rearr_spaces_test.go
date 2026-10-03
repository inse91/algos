package _1592_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rearrange-spaces-between-words

func Test(t *testing.T) {
	cases := []struct {
		name string
		text string
		res  string
	}{
		{
			name: "1",
			text: "  this   is  a sentence ",
			res:  "this   is   a   sentence",
		},
		{
			name: "2",
			text: " practice   makes   perfect",
			res:  "practice   makes   perfect ",
		},
		{
			name: "4",
			text: "a",
			res:  "a",
		},
		{
			name: "5",
			text: " a",
			res:  "a ",
		},
		{
			name: "6",
			text: " a ",
			res:  "a  ",
		},
		{
			name: "7",
			text: "a ",
			res:  "a ",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, reorderSpaces(
				c.text,
			))
		})
	}
}

func reorderSpaces(text string) string {
	var (
		sb          strings.Builder
		parts       []string
		spacesCount int
	)

	for _, r := range text {
		if !unicode.IsSpace(r) {
			sb.WriteRune(r)
			continue
		}

		spacesCount += 1
		if sb.Len() == 0 {
			continue
		}

		parts = append(parts, sb.String())
		sb.Reset()
	}

	if sb.Len() > 0 {
		parts = append(parts, sb.String())
	}

	var (
		sep   string
		extra = strings.Repeat(" ", spacesCount)
	)
	if len(parts) > 1 {
		sep = strings.Repeat(" ", spacesCount/(len(parts)-1))
		extra = strings.Repeat(" ", spacesCount%(len(parts)-1))
	}

	return strings.Join(parts, sep) + extra
}
