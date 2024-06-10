package _844_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/backspace-string-compare/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		t    string
		res  bool
	}{
		{
			name: "1",
			s:    "ab#c",
			t:    "ad#c",
			res:  true,
		},
		{
			name: "2",
			s:    "ab##",
			t:    "c#d#",
			res:  true,
		},
		{
			name: "3",
			s:    "a#c",
			t:    "b",
			res:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, backspaceCompare(c.s, c.t))
		})
	}
}

type msb struct {
	buf []byte
}

func (sb *msb) Init(len int) {
	sb.buf = make([]byte, 0, len)
}

func (sb *msb) Write(b rune) {
	sb.buf = append(sb.buf, byte(b))
}

func (sb *msb) Trim() {
	if len(sb.buf) == 0 {
		return
	}
	sb.buf = sb.buf[:len(sb.buf)-1]
}

func (sb *msb) String() string {
	return string(sb.buf)
}

func backspaceCompare(s string, t string) bool {
	s = grep(s)
	t = grep(t)

	return s == t
}

func grep(s string) string {
	const backspace = '#'

	var ns msb
	ns.Init(len(s))

	for _, r := range s {
		if r == backspace {
			ns.Trim()
			continue
		}
		ns.Write(r)
	}

	return ns.String()
}
