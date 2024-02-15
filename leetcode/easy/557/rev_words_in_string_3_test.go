package _557_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"strings"
	"testing"
)

func TestRevWordsInString(t *testing.T) {
	cases := []struct {
		name string
		in   string
		res  string
	}{
		{
			name: "1",
			in:   "Let's take LeetCode contest",
			res:  "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name: "2",
			in:   "Mr Ding",
			res:  "rM gniD",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := reverseWords(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func reverseWords(s string) string {
	sl := strings.Split(s, " ")
	for i, str := range sl {
		bts := []byte(str)
		slices.Reverse(bts)
		sl[i] = string(bts)
	}

	return strings.Join(sl, " ")
}
