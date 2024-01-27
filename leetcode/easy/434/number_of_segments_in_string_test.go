package _434_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNumberOfSegmentsInString(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "Hello, my name is John",
			res:  5,
		},
		{
			name: "2",
			s:    "Hello",
			res:  1,
		},
		{
			name: "3",
			s:    "",
			res:  0,
		},
		{
			name: "4",
			s:    "      ",
			res:  0,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := countSegments(c.s)
			assert.Equal(t, c.res, res)
		})
	}
}

const space = " "
const empty = ""

func countSegments(s string) int {
	if s == "" {
		return 0
	}
	segs := strings.Split(s, space)
	var res int
	for _, seg := range segs {
		if seg == empty {
			continue
		}
		res++
	}
	return res
}
