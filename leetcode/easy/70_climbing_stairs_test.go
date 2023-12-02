package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test70(t *testing.T) {
	testCases := []struct {
		num  int
		want int
		name string
	}{
		{name: "1", num: 2, want: 2},
		{name: "2", num: 3, want: 3},
		{name: "3", num: 4, want: 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, climbStairs(tc.num))
		})
	}
}

func climbStairs(n int) int {
	a, b := 1, 1
	for n > 1 {
		a, b = b, a+b
		n--
	}
	return b
}
