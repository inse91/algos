package _1556_test

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/slowest-key/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  string
	}{
		{
			name: "1",
			n:    987,
			res:  "987",
		},
		{
			name: "2",
			n:    1234,
			res:  "1.234",
		},
		{
			name: "33",
			n:    21434234,
			res:  "21.434.234",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, thousandSeparator(c.n))
		})
	}
}

func thousandSeparator(n int) string {
	s := strconv.Itoa(n)
	bts := []byte(s)
	slices.Reverse(bts)

	var sb strings.Builder
	for i, b := range bts {
		if i%3 == 0 && i != 0 {
			sb.WriteByte('.')
		}
		sb.WriteByte(b)
	}

	bts = []byte(sb.String())
	slices.Reverse(bts)

	return string(bts)
}
