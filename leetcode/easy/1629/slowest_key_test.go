package _1629_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/slowest-key/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		rt   []int
		keys string
		res  byte
	}{
		{
			name: "1",
			rt:   []int{9, 29, 49, 50},
			keys: "cbcd",
			res:  'c',
		},
		{
			name: "2",
			rt:   []int{12, 23, 36, 46, 62},
			keys: "spuda",
			res:  'a',
		},
		{
			name: "34",
			rt:   []int{23, 34, 43, 59, 62, 80, 83, 92, 97},
			keys: "qgkzzihfc",
			res:  'q',
		},
		{
			name: "43",
			rt:   []int{28, 65, 97},
			keys: "gaf",
			res:  'a',
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, slowestKey(c.rt, c.keys))
		})
	}
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxDur := releaseTimes[0]
	char := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		dur := releaseTimes[i] - releaseTimes[i-1]
		if dur < maxDur {
			continue
		}

		currentKey := keysPressed[i]
		if dur == maxDur {
			char = max(char, currentKey)
			continue
		}

		// dur > maxDur
		maxDur = dur
		char = currentKey
	}

	return char
}
