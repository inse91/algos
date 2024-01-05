package _344_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
		res  []byte
	}{
		{
			name: "1",
			in:   []byte("hello"),
			res:  []byte("olleh"),
		},
		{
			name: "2",
			in:   []byte("a"),
			res:  []byte("a"),
		},
		{
			name: "3",
			in:   []byte("ab"),
			res:  []byte("ba"),
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			reverseString(c.in)
			assert.Equal(t, c.res, c.in)
		})
	}
}

func reverseString(s []byte) {
	slices.Reverse(s)
}
