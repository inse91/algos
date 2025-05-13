package _3498_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-degree-of-a-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "abc",
			res:  148,
		},
		{
			name: "2",
			s:    "zaza",
			res:  160,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, reverseDegree(c.s))
		})
	}
}

func reverseDegree(s string) int {
	var res int
	for i, r := range s {
		res += int(123-r) * (i + 1)
	}

	return res
}
