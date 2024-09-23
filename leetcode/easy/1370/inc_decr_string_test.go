package _1370_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/increasing-decreasing-string/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "aaaabbbbcccc",
			res:  "abccbaabccba",
		},
		{
			name: "2",
			s:    "rat",
			res:  "art",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sortString(c.s))
		})
	}
}

func sortString(s string) string {
	const offset = 97
	var letters [26]int

	fmt.Println(int('a'))
	mx := -1
	var idx int
	for _, r := range s {
		idx = int(r) - offset
		letters[idx]++
		mx = max(mx, letters[idx])
	}

	needOneMore := mx%2 == 1

	var sb strings.Builder
	for i := 0; i < mx/2; i++ {
		// to the right
		for letter, c := range letters {
			if c == 0 {
				continue
			}

			sb.WriteByte(byte(letter + offset))
			letters[letter]--
		}

		// to the right
		var c int
		for letter := len(letters) - 1; letter >= 0; letter-- {
			c = letters[letter]
			if c == 0 {
				continue
			}

			sb.WriteByte(byte(letter + offset))
			letters[letter]--
		}
	}

	if needOneMore {
		// to the right
		for letter, c := range letters {
			if c == 0 {
				continue
			}

			sb.WriteByte(byte(letter + offset))
			letters[letter]--
		}
	}

	return sb.String()
}
