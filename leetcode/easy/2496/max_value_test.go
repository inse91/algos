package _2496_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-value-of-a-string-in-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		strs []string
		res  int
	}{
		{
			name: "1",
			strs: []string{"alic3", "bob", "3", "4", "00000"},
			res:  5,
		},
		{
			name: "2",
			strs: []string{"1", "01", "001", "0001"},
			res:  1,
		},
		{
			name: "3",
			strs: []string{"1", "01", "001", "0002"},
			res:  2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maximumValue(c.strs))
		})
	}
}

func maximumValue(strs []string) int {
	var res int
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err == nil {
			res = max(num, res)
			continue
		}

		res = max(res, len(s))
	}

	return res
}
