package _1502_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/consecutive-characters/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		res  bool
	}{
		{
			name: "1",
			arr:  []int{3, 5, 1},
			res:  true,
		},
		{
			name: "2",
			arr:  []int{1, 2, 4},
			res:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, canMakeArithmeticProgression(c.arr))
		})
	}
}

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 2 {
		return len(arr) == 1
	}

	slices.Sort(arr)
	diff := arr[1] - arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == diff {
			continue
		}

		return false
	}

	return true
}
