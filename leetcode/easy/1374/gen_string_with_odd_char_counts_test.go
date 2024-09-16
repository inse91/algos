package _1374_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
	}{
		{
			name: "1",
			num:  4,
		},
		{
			name: "2",
			num:  2,
		},
		{
			name: "3",
			num:  7,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := generateTheString(c.num)
			assert.True(t, check(res))
		})
	}
}

func generateTheString(n int) string {
	s := strings.Repeat("a", n)
	if n%2 == 1 {
		return s
	}

	return s[:len(s)-1] + "b"
}

func check(s string) bool {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	for _, c := range m {
		if c%2 == 0 {
			return false
		}
	}

	return true
}
