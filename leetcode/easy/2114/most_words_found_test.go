package _2652_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences

func Test(t *testing.T) {
	cases := []struct {
		name      string
		sentences []string
		res       int
	}{
		{
			name: "1",
			sentences: []string{
				"alice and bob love leetcode",
				"i think so too",
				"this is great thanks very much",
			},
			res: 6,
		},
		{
			name: "2",
			sentences: []string{
				"please wait",
				"continue to fight",
				"continue to win",
			},
			res: 3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, mostWordsFound(c.sentences))
		})
	}
}

func mostWordsFound(sentences []string) int {
	var res, c int
	for _, s := range sentences {
		c = 0
		for _, r := range s {
			if r == ' ' {
				c += 1
			}
		}
		if c > 0 {
			c += 1
		}

		res = max(res, c)
	}

	return res
}
