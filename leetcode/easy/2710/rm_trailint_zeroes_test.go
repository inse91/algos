package _2651_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-trailing-zeros-from-a-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  string
		res  string
	}{
		{
			name: "1",
			num:  "51230100",
			res:  "512301",
		},
		{
			name: "2",
			num:  "123",
			res:  "123",
		},
		{
			name: "3",
			num:  "000",
			res:  "",
		},
		{
			name: "743",
			num:  "9",
			res:  "9",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, removeTrailingZeros(c.num))
		})
	}
}

func removeTrailingZeros(num string) string {
	idx := len(num) - 1
	hasOnlyZeroes := true
	for i := idx; i >= 0; i-- {
		if num[i] == '0' {
			continue
		}

		idx = i
		hasOnlyZeroes = false
		break
	}

	if hasOnlyZeroes {
		return ""
	}

	return num[:idx+1]
}
