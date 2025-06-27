package _2053_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kth-distinct-string-in-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []string
		k    int
		res  string
	}{
		{
			name: "1",
			arr:  []string{"d", "b", "c", "b", "c", "a"},
			k:    2,
			res:  "a",
		},
		{
			name: "2",
			arr:  []string{"aaa", "aa", "a"},
			k:    1,
			res:  "aaa",
		},
		{
			name: "3",
			arr:  []string{"a", "b", "a"},
			k:    3,
			res:  "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, kthDistinct(c.arr, c.k))
		})
	}
}

func kthDistinct(arr []string, k int) string {
	var c int
	set := make(map[string]bool, k)
	for _, s := range arr {
		_, ok := set[s]
		if !ok {
			set[s] = true
			c++

			continue
		}

		set[s] = false
		c--
	}

	if c < k {
		return ""
	}

	for _, s := range arr {
		flag, ok := set[s]
		if !ok || !flag {
			continue
		}

		if k == 1 {
			return s
		}

		k--
	}

	return ""
}
