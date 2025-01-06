package _1662_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/

func Test(t *testing.T) {
	cases := []struct {
		name string
		w1   []string
		w2   []string
		res  bool
	}{
		{
			name: "1",
			w1:   []string{"ab", "c"},
			w2:   []string{"a", "bc"},
			res:  true,
		},
		{
			name: "2",
			w1:   []string{"a", "cb"},
			w2:   []string{"ab", "c"},
			res:  false,
		},
		{
			name: "3",
			w1:   []string{"abc", "d", "defg"},
			w2:   []string{"abcddefg"},
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, arrayStringsAreEqual(c.w1, c.w2))
		})
	}
}

func arrayStringsAreEqual_(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1 := container{store: word1}
	w2 := container{store: word2}

	for {
		ch1, done1 := w1.next()
		ch2, done2 := w2.next()

		if !done1 && !done2 {
			break
		}

		if !done1 || !done2 {
			return false
		}

		if ch1 != ch2 {
			return false
		}
	}

	return true
}

type container struct {
	store       []string
	currentIdx  int
	currentElem int
}

func (c *container) next() (byte, bool) {
	if c.currentElem == len(c.store) {
		return 0, false
	}

	if c.currentIdx == len(c.store[c.currentElem]) {
		c.currentElem++
		c.currentIdx = 0
		return c.next()
	}

	ch := c.store[c.currentElem][c.currentIdx]
	c.currentIdx++

	return ch, true
}
