package _2243_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/calculate-digit-sum-of-a-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		k    int
		res  string
	}{
		{
			name: "1",
			s:    "11111222223",
			k:    3,
			res:  "135",
		},
		{
			name: "2",
			s:    "00000000",
			k:    3,
			res:  "000",
		},
		{
			name: "3",
			s:    "123456789",
			k:    3,
			res:  "126",
		},
		{
			name: "4",
			s:    "123456789",
			k:    4,
			res:  "99",
		},
		{
			name: "5",
			s:    "3465",
			k:    3,
			res:  "135",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, digitSum(c.s, c.k))
		})
	}
}

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}

	var sb strings.Builder
	sb.Grow(len(s) / k)

	var acc int
	for i, r := range s {
		if i%k == 0 && i != 0 {
			char := strconv.Itoa(acc)
			sb.WriteString(char)

			acc = 0
		}

		acc += int(r - '0')
		if i == len(s)-1 {
			sb.WriteString(strconv.Itoa(acc))
		}
	}

	res := sb.String()

	return digitSum(res, k)
}
