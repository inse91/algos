package _2582_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/pass-the-pillow

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		t    int
		res  int
	}{
		{
			name: "1",
			n:    4,
			t:    5,
			res:  2,
		},
		{
			name: "2",
			n:    3,
			t:    2,
			res:  3,
		},
		{
			name: "14",
			n:    3,
			t:    1,
			res:  2,
		},
		{
			name: "17",
			n:    3,
			t:    5,
			res:  2,
		},
		{
			name: "23",
			n:    4,
			t:    11,
			res:  2,
		},
		{
			name: "48",
			n:    8,
			t:    9,
			res:  6,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, passThePillow(c.n, c.t))
		})
	}
}

func passThePillow(n int, time int) int {
	circleLn := (n - 1) * 2
	time = time%circleLn + 1

	if time <= n {
		return time
	}

	return 2*n - time
}
