package _476_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestNumberComplement(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  5,
			res:  2,
		},
		{
			name: "2",
			num:  1,
			res:  0,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findComplement(c.num)
			assert.Equal(t, c.res, res)
		})
	}
}

func findComplement(num int) int {
	s := strconv.FormatInt(int64(num), 2)

	const zero rune = '0'
	const one rune = '1'
	var sb strings.Builder
	sb.Grow(len(s))
	for _, r := range s {
		if r == zero {
			sb.WriteRune(one)
		} else {
			sb.WriteRune(zero)
		}
	}

	i, _ := strconv.ParseInt(sb.String(), 2, 32)
	return int(i)
}
