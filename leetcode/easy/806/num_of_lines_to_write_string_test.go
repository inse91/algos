package _806_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumOfLinesToWriteString(t *testing.T) {
	cases := []struct {
		name string
		w    []int
		s    string
		res  []int
	}{
		{
			name: "1",
			w:    []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			s:    "abcdefghijklmnopqrstuvwxyz",
			res:  []int{3, 60},
		},
		{
			name: "2",
			w:    []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			s:    "bbbcccdddaaa",
			res:  []int{2, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, numberOfLines(c.w, c.s))
		})
	}
}

func numberOfLines(widths []int, s string) []int {
	const gap = 97

	var ttl = 1
	var cur int
	for _, r := range s {
		w := widths[r-gap]
		cur += w

		if cur > 100 {
			cur = w
			ttl++
		}
	}

	return []int{ttl, cur}
}
