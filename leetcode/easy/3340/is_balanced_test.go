package _3340_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-balanced-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  string
		res  bool
	}{
		{
			name: "1",
			num:  "1234",
			res:  false,
		},
		{
			name: "2",
			num:  "24123",
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isBalanced(c.num))
		})
	}
}

func isBalanced(num string) bool {
	var even, odd int
	for i, dig := range num {
		n := int(dig - '0')
		if i%2 == 0 {
			even += n
			continue
		}

		odd += n
	}

	return even == odd
}
