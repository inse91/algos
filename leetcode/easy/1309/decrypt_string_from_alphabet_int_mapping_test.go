package _1309_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "10#11#12",
			res:  "jkab",
		},
		{
			name: "2",
			s:    "1326#",
			res:  "acz",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, freqAlphabets(c.s))
		})
	}
}

func freqAlphabets(s string) string {
	s = reverse(s)
	var sb strings.Builder
	var b byte
	var idx int
	for i := 0; i < len(s); i++ {
		b = s[i]
		if b == '#' {
			var acc int
			for j := 0; j < 2; j++ {
				i++
				b = s[i]
				ss := string(b)
				_ = ss
				acc += digits[s[i]] * pow10(j)
			}
			sb.WriteByte(letters[acc])
			continue
		}

		idx = digits[b]
		sb.WriteByte(letters[idx])
	}

	return reverse(sb.String())
}

var letters = []byte{' ', // to start with idx=1
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r',
	's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

var digits = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func reverse(s string) string {
	bts := []byte(s)
	slices.Reverse(bts)

	return string(bts)
}

func pow10(i int) int {
	switch i {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	default:
		panic("wrong i")
	}
}
