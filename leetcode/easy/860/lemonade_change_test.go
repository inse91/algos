package _860_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/lemonade-change/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		res  bool
	}{
		{
			name: "1",
			in:   []int{5, 5, 5, 10, 20},
			res:  true,
		},
		{
			name: "2",
			in:   []int{5, 5, 10, 10, 20},
			res:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, lemonadeChange(c.in))
		})
	}
}

func lemonadeChange(bills []int) bool {
	m := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}

	for _, b := range bills {
		if b == 5 {
			m[5]++
			continue
		}

		if b == 10 {
			if m[5] == 0 {
				return false
			}
			m[10]++
			m[5]--
			continue
		}

		if b == 20 {
			if m[5] == 0 {
				return false
			}
			m[5]--
			if m[10] == 0 {
				if m[5] < 2 {
					return false
				}
				m[5] -= 2
				continue
			}

			m[10]--
		}
	}

	return true
}
