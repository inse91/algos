package _191

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestNumberOf1bits(t *testing.T) {
	testCases := []struct {
		name string
		in   uint32
		res  int
	}{
		{
			name: "1",
			in:   0b00000000000000000000000000001011,
			res:  3,
		},
		{
			name: "2",
			in:   0b00000000000000000000000010000000,
			res:  1,
		},
		{
			name: "3",
			in:   0b11111111111111111111111111111101,
			res:  31,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := hammingWeight(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}

func hammingWeight(num uint32) int {

	f := strconv.FormatInt(int64(num), 2)

	return strings.Count(f, "1")
}
