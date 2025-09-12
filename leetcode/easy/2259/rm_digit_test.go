package _2259_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-digit-from-number-to-maximize-result

func Test(t *testing.T) {
	cases := []struct {
		name  string
		num   string
		digit byte
		res   string
	}{
		{
			name:  "1",
			num:   "123",
			digit: '3',
			res:   "12",
		},
		{
			name:  "2",
			num:   "1231",
			digit: '1',
			res:   "231",
		},
		{
			name:  "3",
			num:   "551",
			digit: '5',
			res:   "51",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, removeDigit(c.num, c.digit))
		})
	}
}

func removeDigit(number string, digit byte) string {
	var res string
	for i := len(number) - 1; i >= 0; i-- {
		if number[i] == digit {
			res = max(res, number[:i]+number[i+1:])
		}
	}

	return res
}
