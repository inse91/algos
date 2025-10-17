package _1417_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reformat-the-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		cost string
		res  string
	}{
		{
			name: "1",
			cost: "a0b1c2",
			res:  "b0a1c2",
		},
		{
			name: "2",
			cost: "leetcode",
			res:  "",
		},
		{
			name: "3",
			cost: "1229857369",
			res:  "",
		},
		{
			name: "11",
			cost: "a0b1c222",
			res:  "",
		},
		{
			name: "17",
			cost: "j",
			res:  "j",
		},
		{
			name: "23",
			cost: "covid2019",
			res:  "c2o0v1i9d",
		},
		{
			name: "23_1",
			cost: "cov2019",
			res:  "9c2o0v1",
		},
		{
			name: "27",
			cost: "a0",
			res:  "0a",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, reformat(c.cost))
		})
	}
}

func reformat(s string) string {
	digs := make([]rune, 0, len(s)/3)
	chars := make([]rune, 0, len(s)*2/3)

	for _, r := range s {
		if unicode.IsDigit(r) {
			digs = append(digs, r)
			continue
		}

		chars = append(chars, r)
	}

	if len(chars) == 0 || len(digs) == 0 {
		if len(s) <= 1 {
			return s
		}

		return ""
	}

	diff := len(chars) - len(digs)
	if mod(diff) > 1 {
		return ""
	}

	var (
		sb           strings.Builder
		digsIsLonger bool = len(digs) > len(chars)
	)

	if digsIsLonger {
		sb.WriteRune(digs[len(digs)-1])
	}
	for i := range chars {
		sb.WriteRune(chars[i])
		if i < len(digs) {
			sb.WriteRune(digs[i])
			continue
		}
	}

	newS := sb.String()
	if newS != s {
		return newS
	}

	switch len(s) {
	case 1:
		return s
	case 2:
		return string(s[1]) + string(s[0])
	default:
	}

	return string(s[2]) + string(s[1]) + string(s[0]) + s[3:]
}

func mod(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
