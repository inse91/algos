package _1652_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/defuse-the-bomb/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		code []int
		k    int
		res  []int
	}{
		{
			name: "1",
			code: []int{5, 7, 1, 4},
			k:    3,
			res:  []int{12, 10, 16, 13},
		},
		{
			name: "2",
			code: []int{1, 2, 3, 4},
			k:    0,
			res:  []int{0, 0, 0, 0},
		},
		{
			name: "3",
			code: []int{2, 4, 9, 3},
			k:    -2,
			res:  []int{12, 5, 6, 13},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, decrypt(c.code, c.k))
		})
	}
}

func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	if k == 0 {
		return res
	}

	var (
		init        int
		needReverse bool
	)
	if k < 0 {
		k = -k
		needReverse = true
		slices.Reverse(code)
	}

	for i, j := 0, 1; i < k; i++ {
		init += code[j]
		j++
	}

	endGetter := func(start int) int {
		end := start + k
		end = end % len(code)

		return code[end]
	}

	res[0] = init
	for i := 1; i < len(res); i++ {
		start := code[i]
		end := endGetter(i)
		res[i] = init - start + end
		init = res[i]
	}

	if needReverse {
		slices.Reverse(res)
	}

	return res
}
