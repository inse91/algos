package _2283_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  string
		res  bool
	}{
		{
			name: "1",
			num:  "1210",
			res:  true,
		},
		{
			name: "2",
			num:  "030",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, digitCount(c.num))
		})
	}
}

func digitCount(num string) bool {
	m := make(map[int]int)
	for i, digit := range num {
		_ = i
		m[int(digit-'0')]++
	}

	for i, digit := range num {
		v := m[i]
		n := int(digit - '0')
		if v != n {
			return false
		}
	}

	return true
}
