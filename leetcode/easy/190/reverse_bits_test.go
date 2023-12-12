package _190

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestReverseBits(t *testing.T) {
	testCases := []struct {
		name string
		in   uint32
		res  uint32
	}{
		{
			name: "1",
			in:   43261596,
			res:  964176192,
		},
		{
			name: "2",
			in:   0,
			res:  0,
		},
		{
			name: "3",
			in:   1,
			res:  2147483648,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := reverseBits(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}

}

func reverseBits(num uint32) uint32 {
	f := strconv.FormatInt(int64(num), 2)

	sb := strings.Builder{}
	sb.Grow(32)
	sb.WriteString(strings.Repeat("0", 32-len(f)))
	sb.WriteString(f)

	f = stringsReverse(sb.String())
	number, err := strconv.ParseInt(f, 2, 32)
	if err != nil {
		panic(err)
	}
	return uint32(number)

}

func stringsReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
