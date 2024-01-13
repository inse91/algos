package _405__test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"slices"
	"strings"
	"testing"
)

func TestConvertToHex(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  string
	}{
		{
			name: "1",
			in:   26,
			res:  "1a",
		},
		{
			name: "2",
			in:   0,
			res:  "0",
		},
		{
			name: "3",
			in:   1,
			res:  "1",
		},
		{
			name: "4",
			in:   255,
			res:  "ff",
		},
		{
			name: "5",
			in:   303,
			res:  "12f",
		},
		//{
		//	name: "5",
		//	in:   -6,
		//	res:  "",
		//},
		{
			name: "6",
			in:   -1,
			res:  "ffffffff",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.res, toHex(c.in))
		})
	}
}

var pow32of2 = int(math.Pow(float64(2), float64(32)))

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	if num < 0 {
		num = num + pow32of2
	}

	chars := [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	var sb strings.Builder

	for num != 0 {
		n := num % 16
		sb.WriteByte(chars[n])
		fmt.Println(sb.String())
		num /= 16
	}

	bts := []byte(sb.String())
	slices.Reverse(bts)

	return string(bts)
}
