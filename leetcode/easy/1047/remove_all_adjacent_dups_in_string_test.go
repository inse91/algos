package _1047_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "abbaca",
			res:  "ca",
		},
		{
			name: "1_1",
			s:    "abbbaca",
			res:  "abaca",
		},
		{
			name: "1_2",
			s:    "abba",
			res:  "",
		},
		{
			name: "2",
			s:    "azxxzy",
			res:  "ay",
		},
		{
			name: "2_1",
			s:    "azzy",
			res:  "ay",
		},
		{
			name: "3",
			s:    "acbb",
			res:  "ac",
		},
		{
			name: "4",
			s:    "acbd",
			res:  "acbd",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, removeDuplicates(c.s))
		})
	}
}

type stack struct {
	bytes []byte
}

func newStack(l int) stack {
	return stack{bytes: make([]byte, 0, l)}
}

func (st *stack) push(b byte) {
	st.bytes = append(st.bytes, b)
}

func (st *stack) peek() byte {
	if len(st.bytes) == 0 {
		return 0
	}
	return st.bytes[len(st.bytes)-1]
}

func (st *stack) pop() {
	st.bytes = st.bytes[:st.len()-1]
}

func (st *stack) len() int {
	return len(st.bytes)
}

func (st *stack) string() string {
	return string(st.bytes)
}

func removeDuplicates(s string) string {
	st := newStack(len(s))
	st.push(s[0])

	for i := range s[:len(s)-1] {
		p := st.peek()
		b := s[i+1]
		if p == b {
			st.pop()
			continue
		}

		st.push(b)
	}

	return st.string()
}
