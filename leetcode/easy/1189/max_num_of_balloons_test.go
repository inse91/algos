package _1189_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "nlaebolko",
			res:  1,
		},
		{
			name: "2",
			s:    "loonbalxballpoon",
			res:  2,
		},
		{
			name: "3",
			s:    "leetcode",
			res:  0,
		},
		{
			name: "17",
			s:    "balon",
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxNumberOfBalloons(c.s))
		})
	}
}

const (
	b = 'b'
	a = 'a'
	l = 'l'
	o = 'o'
	n = 'n'
)

func maxNumberOfBalloons(text string) int {
	m := map[rune]int{
		b: 0, a: 0, l: 0, o: 0, n: 0,
	}
	for _, r := range text {
		switch r {
		case b:
			m[b]++
		case a:
			m[a]++
		case l:
			m[l]++
		case o:
			m[o]++
		case n:
			m[n]++
		}
	}

	m[o] /= 2
	m[l] /= 2

	arr := make([]int, 0, len(m))
	for _, v := range m {
		arr = append(arr, v)
	}

	if len(arr) == 0 {
		return 0
	}

	minn := arr[0]
	for _, nn := range arr[1:] {
		minn = min(minn, nn)
	}

	return minn
}
