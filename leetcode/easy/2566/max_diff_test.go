package _2566_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-difference-by-remapping-a-digit

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  11891,
			res:  99009,
		},
		{
			name: "2",
			num:  90,
			res:  99,
		},
		{
			name: "3",
			num:  37398834,
			res:  90900090,
		},
		{
			name: "22",
			num:  91,
			res:  98,
		},
		{
			name: "27",
			num:  393,
			res:  909,
		},
		{
			name: "31",
			num:  965,
			res:  930,
		},
		{
			name: "194",
			num:  900,
			res:  999,
		},
		{
			name: "200",
			num:  95334529,
			res:  94000409,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minMaxDifference(c.num))
		})
	}
}

func minMaxDifference(num int) int {
	s := strconv.Itoa(num)
	var (
		first = rune(s[0])
		res   int
	)

	for i, r := range s {
		if r != first {
			continue
		}

		dig := 9
		for range len(s) - i - 1 {
			dig *= 10
		}

		res += dig
	}

	if first != '9' {
		return res
	}

	var second rune = -1
	for i, r := range s {
		if r == first {
			continue
		}

		if second != -1 && r != second {
			continue
		}

		second = r
		dig := 9 - int(r-'0')
		for range len(s) - i - 1 {
			dig *= 10
		}

		res += dig
	}

	return res
}
