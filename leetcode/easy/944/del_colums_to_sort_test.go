package _944_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/delete-columns-to-make-sorted/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		strs []string
		res  int
	}{
		{
			name: "1",
			strs: []string{"cba", "daf", "ghi"},
			res:  1,
		},
		{
			name: "2",
			strs: []string{"a", "b"},
			res:  0,
		},
		{
			name: "3",
			strs: []string{"zyx", "wvu", "tsr"},
			res:  3,
		},
		{
			name: "6",
			strs: []string{"rrjk", "furt", "guzm"},
			res:  2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minDeletionSize(c.strs))
		})
	}
}

func minDeletionSize(strs []string) int {
	if len(strs) <= 1 {
		return 0
	}

	deleted := make(map[int]struct{})

	for i := 1; i < len(strs); i++ {
		cur := strs[i]
		prev := strs[i-1]

		for idx, r := range cur {
			if byte(r) < prev[idx] {
				deleted[idx] = struct{}{}
			}
		}
	}

	return len(deleted)
}
