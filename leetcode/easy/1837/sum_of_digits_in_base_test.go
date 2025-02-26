package _1837_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-digits-in-base-k/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		k    int
		res  int
	}{
		{
			name: "1",
			n:    34,
			k:    6,
			res:  9,
		},
		{
			name: "2",
			n:    10,
			k:    10,
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumBase(c.n, c.k))
		})
	}
}

func sumBase(n int, k int) int {
	var res int
	s := strconv.FormatInt(int64(n), k)
	for _, r := range s {
		num, err := strconv.ParseInt(string(r), 10, 8)
		if err != nil {
			panic(err)
		}

		res += int(num)
	}

	return res
}
