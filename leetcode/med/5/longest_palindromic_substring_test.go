package _5_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	cases := []struct {
		name string
		in   string
		res  string
	}{
		{
			name: "1",
			in:   "babad",
			res:  "bab",
		},
		{
			name: "2_1",
			in:   "cbbd",
			res:  "bb",
		},
		{
			name: "2_2",
			in:   "acbbd",
			res:  "bb",
		},
		{
			name: "2_3",
			in:   "acbbbd",
			res:  "bbb",
		},
		{
			name: "2_4",
			in:   "acbbbcd",
			res:  "cbbbc",
		},
		{
			name: "3",
			in:   "a",
			res:  "a",
		},
		{
			name: "4",
			in:   "ac",
			res:  "a",
		},
		{
			name: "5",
			in:   "aba",
			res:  "aba",
		},
		{
			name: "5_1",
			in:   "abba",
			res:  "abba",
		},
		{
			name: "6",
			in:   "abacd",
			res:  "aba",
		},
		{
			name: "7",
			in:   "bb",
			res:  "bb",
		},
		{
			name: "7_1",
			in:   "bbb",
			res:  "bbb",
		},
		{
			name: "7_1_1",
			in:   "bbbb",
			res:  "bbbb",
		},
		{
			name: "7_1_2",
			in:   "abbbba",
			res:  "abbbba",
		},
		{
			name: "7_2",
			in:   "bbbbb",
			res:  "bbbbb",
		},
		{
			name: "7_3",
			in:   "bbbbbb",
			res:  "bbbbbb",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			//t.Parallel()
			res := longestPalindrome(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	lens := len(s)

	var longest string = ""
	var curLongest string

	var prevIdx int
	var nextIdx int
	for i := range s {
		// skip last letter
		if i == lens-1 {
			//if len(longest) >= lens-i {
			break
		}

		// "aba" case
		if s[i+1] != s[i] || (i > 0 && s[i] == s[i-1]) {
			curLongest = string(s[i])
			prevIdx = i - 1
			nextIdx = i + 1
			for {
				if prevIdx < 0 || nextIdx > lens-1 {
					break
				}
				if s[prevIdx] != s[nextIdx] {
					break
				}
				curLongest = string(s[prevIdx]) + curLongest + string(s[nextIdx])
				prevIdx--
				nextIdx++
			}
			if len(longest) < len(curLongest) {
				longest = curLongest
			}
			//continue
		}

		// "abba", "bb" case
		if s[i+1] == s[i] {
			curLongest = string(s[i]) + string(s[i+1])
			prevIdx = i - 1
			nextIdx = i + 2
			for {
				if prevIdx < 0 || nextIdx > lens-1 {
					break
				}
				if s[prevIdx] != s[nextIdx] {
					break
				}
				curLongest = string(s[prevIdx]) + curLongest + string(s[nextIdx])
				prevIdx--
				nextIdx++
			}
			if len(longest) < len(curLongest) {
				longest = curLongest
			}
		}
	}
	return longest
}
