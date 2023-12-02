package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun12(t *testing.T) {
	testCases := []struct {
		num  int
		want string
		name string
	}{
		{name: "1", num: 3, want: "III"},
		{name: "2", num: 58, want: "LVIII"},
		{name: "3", num: 1994, want: "MCMXCIV"},
		//{num: 3, want: "III"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := intToRoman(tc.num)
			assert.Equal(t, tc.want, got)
		})
	}
}
