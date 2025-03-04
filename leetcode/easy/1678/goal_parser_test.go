package _1678_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/goal-parser-interpretation/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		c    string
		res  string
	}{
		{
			name: "1",
			c:    "G()(al)",
			res:  "Goal",
		},
		{
			name: "2",
			c:    "G()()()()(al)",
			res:  "Gooooal",
		},
		{
			name: "3",
			c:    "(al)G(al)()()G",
			res:  "alGalooG",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, interpret(c.c))
		})
	}
}

func interpret(command string) string {
	var sb strings.Builder
	for i, r := range command {
		if r == 'G' {
			sb.WriteByte('G')
			continue
		}

		if r != '(' {
			continue
		}

		// r == '('
		nextByte := command[i+1]
		if nextByte == ')' {
			sb.WriteByte('o')
			continue
		}

		sb.WriteString("al")
	}

	return sb.String()
}
