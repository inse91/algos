package _205

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsomorphicStrings(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		t    string
		res  bool
	}{
		{
			name: "1",
			s:    "egg",
			t:    "add",
			res:  true,
		},
		{
			name: "2",
			s:    "foo",
			t:    "bar",
			res:  false,
		},
		{
			name: "3",
			s:    "paper",
			t:    "title",
			res:  true,
		},
		{
			name: "4",
			s:    "badc",
			t:    "baba",
			res:  false,
		},
		{
			name: "5",
			s:    "bada",
			t:    "babc",
			res:  false,
		},
		{
			name: "6",
			s:    "abc",
			t:    "def",
			res:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isIsomorphic(tc.s, tc.t)
			assert.Equal(t, tc.res, res)
		})
	}
}

func isIsomorphic(s string, t string) bool {

	m := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := range s {
		b := s[i]
		b2 := t[i]

		v, ok := m[b]
		if ok && v != t[i] {
			return false
		}

		v, ok = m2[b2]
		if ok && v != s[i] {
			return false
		}

		m[b] = t[i]
		m2[b2] = s[i]
	}

	return true
}
