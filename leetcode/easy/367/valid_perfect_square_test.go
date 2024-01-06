package _367_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPerfectSquare(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  bool
	}{
		{
			name: "1",
			in:   16,
			res:  true,
		},
		{
			name: "2",
			in:   14,
			res:  false,
		},
		{
			name: "3",
			in:   1,
			res:  true,
		},
		{
			name: "4",
			in:   4,
			res:  true,
		},
		{
			name: "5",
			in:   9,
			res:  true,
		},
		{
			name: "6",
			in:   7,
			res:  false,
		},
		{
			name: "7",
			in:   13,
			res:  false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := isPerfectSquare(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func isPerfectSquare(num int) bool {
	switch num - num/10*10 {
	case 2, 3, 7:
		return false
	default:
	}
	var st int
	//for st = 0; st*st <= num; st++ {
	for {
		sq := st * st
		if sq < num {
			st++
			continue
		}
		return st*st == num
	}
}
