package _696_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBinarySubstrings(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "00110011",
			res:  6,
		},
		{
			name: "2",
			s:    "10101",
			res:  4,
		},
		{
			name: "3",
			s:    "00001111",
			res:  4,
		},
		{
			name: "4",
			s:    "000000",
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := countBinarySubstrings(c.s)
			assert.Equal(t, c.res, res)
		})
	}
}

func countBinarySubstrings(s string) (res int) {
	hist := []int{0, 1}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			hist[len(hist)-1]++
			continue
		}
		res += min(hist[len(hist)-1], hist[len(hist)-2])
		hist[len(hist)-2] = hist[len(hist)-1]
		hist[len(hist)-1] = 1
		//hist = append(hist, 1)
	}
	res += min(hist[len(hist)-1], hist[len(hist)-2])

	return res
}
