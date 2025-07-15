package _3325_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/decode-the-message

func Test(t *testing.T) {
	cases := []struct {
		name string
		key  string
		msg  string
		res  string
	}{
		{
			name: "1",
			key:  "the quick brown fox jumps over the lazy dog",
			msg:  "vkbs bs t suepuv",
			res:  "this is a secret",
		},
		{
			name: "2",
			key:  "eljuxhpwnyrdgtqkviszcfmabo",
			msg:  "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
			res:  "the five boxing wizards jump quickly",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, decodeMessage(c.key, c.msg))
		})
	}
}

func decodeMessage(key string, message string) string {
	m := make(map[rune]rune, 26)
	var c int
	for _, r := range key {
		if _, ok := m[r]; ok || r == ' ' {
			continue
		}

		m[r] = rune(c + 'a')
		c++
	}

	var sb strings.Builder
	for _, r := range message {
		if r == ' ' {
			sb.WriteRune(' ')
			continue
		}

		sb.WriteRune(m[r])
	}

	return sb.String()
}
