package _1323_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

// https://leetcode.com/problems/maximum-69-number/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    9669,
			res:  9969,
		},
		{
			name: "2",
			n:    9996,
			res:  9999,
		},
		{
			name: "3",
			n:    9999,
			res:  9999,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maximum69Number(c.n))
		})
	}

}

func maximum69Number(num int) int {
	s := strconv.Itoa(num)

	var ns strings.Builder
	var isAlreadyChanged bool
	for _, r := range s {
		if r == '6' && !isAlreadyChanged {
			ns.WriteRune('9')
			isAlreadyChanged = true
			continue
		}

		ns.WriteRune(r)
	}

	n, err := strconv.Atoi(ns.String())
	if err != nil {
		panic(err)
	}

	return n
}
