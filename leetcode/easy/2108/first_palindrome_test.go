package _2108_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-first-palindromic-string-in-the-array

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   string
	}{
		{
			name:  "1",
			words: []string{"abc", "car", "ada", "racecar", "cool"},
			res:   "ada",
		},
		{
			name:  "2",
			words: []string{"def", "ghi"},
			res:   "",
		},
		{
			name:  "3",
			words: []string{"notapalindrome", "racecar"},
			res:   "racecar",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, firstPalindrome(c.words))
		})
	}
}

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}

	return ""
}

func isPalindrome(s string) bool {
	for i := range s[:len(s)/2] {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
