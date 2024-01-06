package _389_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDifference(t *testing.T) {
	cases := []struct {
		name string
		in1  string
		in2  string
		res  byte
	}{
		{
			name: "1",
			in1:  "abcd",
			in2:  "abcde",
			res:  byte('e'),
		},
		//{
		//	name: "2",
		//	in1:  "a",
		//	in2:  "b",
		//	res:  byte('e'),
		//},
		//{
		//	name: "3",
		//	in1:  "aa",
		//	in2:  "ab",
		//	res:  "a",
		//},
		{
			name: "4",
			in1:  "a",
			in2:  "aa",
			res:  byte('a'),
		},
		{
			name: "5",
			in1:  "ae",
			in2:  "aea",
			res:  byte('a'),
		},
		//{
		//	name: "6",
		//	in1:  "a",
		//	in2:  "a",
		//	res:  "",
		//},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findTheDifference(c.in1, c.in2)
			assert.Equal(t, c.res, res)
		})
	}
}

func findTheDifference(s string, t string) byte {
	var chars [26]int

	for _, r := range s {
		chars[r-'a']++
	}

	for _, r := range t {
		chars[r-'a']--
		if chars[r-'a'] < 0 {
			return byte(r)
		}
	}
	return 0
}
