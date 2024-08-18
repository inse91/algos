package _680_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "aba",
			res:  true,
		},
		{
			name: "2",
			s:    "abca",
			res:  true,
		},
		{
			name: "3",
			s:    "abc",
			res:  false,
		},
		{
			name: "4",
			s:    "deddde",
			res:  true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := validPalindrome(c.s)
			assert.Equal(t, c.res, res)
		})
	}
}

func validPalindrome(s string) bool {
	if isPal(s) {
		return true
	}

	for i := range s {
		bts := []byte(s)
		bts = slices.Delete(bts, i, i+1)

		newS := string(bts)
		if isPal(newS) {
			return true
		}
	}

	return false
}

func isPal(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i := range s {
		l := len(s)
		if i >= l/2+1 {
			break
		}
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
