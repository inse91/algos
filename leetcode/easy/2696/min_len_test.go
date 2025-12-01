package _2520_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-string-length-after-removing-substrings

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "ABFCACDB",
			res:  2,
		},
		{
			name: "2",
			s:    "ACBBD",
			res:  5,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minLength(c.s))
		})
	}
}

func minLength(s string) int {
	rpl := strings.NewReplacer("AB", "", "CD", "")
	for {
		ln := len(s)
		s = rpl.Replace(s)
		if ln == len(s) {
			break
		}
	}

	return len(s)
}
