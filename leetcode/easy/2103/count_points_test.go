package _2103_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rings-and-rods

func Test(t *testing.T) {
	cases := []struct {
		name  string
		rings string
		res   int
	}{
		{
			name:  "1",
			rings: "B0B6G0R6R0R6G9",
			res:   1,
		},
		{
			name:  "2",
			rings: "B0R0G0R9R0B0G0",
			res:   1,
		},
		{
			name:  "3",
			rings: "G4",
			res:   0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countPoints(c.rings))
		})
	}
}

func countPoints(rings string) int {
	type rgb struct {
		r, g, b bool
	}

	var m [100]rgb
	for i := 0; i < len(rings); i += 2 {
		rod, err := strconv.Atoi(string(rings[i+1]))
		if err != nil {
			panic(err)
		}

		switch rings[i] {
		case 'R':
			m[rod] = rgb{
				r: true,
				g: m[rod].g,
				b: m[rod].b,
			}
		case 'G':
			m[rod] = rgb{
				r: m[rod].r,
				g: true,
				b: m[rod].b,
			}
		case 'B':
			m[rod] = rgb{
				r: m[rod].r,
				g: m[rod].g,
				b: true,
			}
		}
	}

	var res int
	for _, v := range m {
		if v.r && v.g && v.b {
			res++
		}
	}

	return res
}
