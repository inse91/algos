package _717_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/1-bit-and-2-bit-characters/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		bits []int
		res  bool
	}{
		{
			name: "1",
			bits: []int{1, 0, 0},
			res:  true,
		},
		{
			name: "2",
			bits: []int{1, 1, 1, 0},
			res:  false,
		},
		{
			name: "3",
			bits: []int{1, 1, 0, 0},
			res:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isOneBitCharacter(c.bits))
		})
	}
}

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); {
		if i == len(bits)-1 && bits[i] == 0 {
			return true
		}
		if bits[i] == 0 {
			i++
		} else {
			i += 2
		}
	}

	return false
}
