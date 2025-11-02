package _2917_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-squares-of-special-elements

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			nums: []int{7, 12, 9, 8, 9, 15},
			k:    4,
			res:  9,
		},
		{
			name: "2",
			nums: []int{2, 12, 1, 11, 4, 5},
			k:    6,
			res:  0,
		},
		{
			name: "3",
			nums: []int{10, 8, 5, 9, 11, 6, 8},
			k:    1,
			res:  15,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findKOr(c.nums, c.k))
		})
	}
}

func findKOr(nums []int, k int) int {
	var (
		strs = make([]string, 0, len(nums))
		mx   int
	)
	for _, num := range nums {
		s := strconv.FormatInt(int64(num), 2)
		strs = append(strs, s)
		mx = max(mx, len(s))
	}

	for i, s := range strs {
		if len(s) == mx {
			continue
		}

		diff := mx - len(s)
		strs[i] = strings.Repeat("0", diff) + s
	}

	var sb strings.Builder
	for i := range mx {
		var (
			b byte = '0'
			c int
		)
		for _, s := range strs {
			if s[i] == '1' {
				c++
			}
			if c == k {
				break
			}
		}

		if c >= k {
			b = '1'
		}
		sb.WriteByte(b)
	}

	ktr, err := strconv.ParseInt(sb.String(), 2, 64)
	if err != nil {
		panic(err)
	}

	return int(ktr)
}
