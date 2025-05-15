package _3461_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "3902",
			res:  true,
		},
		{
			name: "2",
			s:    "34789",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, hasSameDigits(c.s))
		})
	}
}

func hasSameDigits(s string) bool {
	array := make([]int, 0, len(s))
	for _, r := range s {
		array = append(array, int(r-'0'))
	}

	for {
		if len(array) == 2 {
			break
		}

		for i := range array[:len(array)-1] {
			array[i] = (array[i] + array[i+1]) % 10
		}
		array = array[:len(array)-1]
	}

	return array[0] == array[1]
}
