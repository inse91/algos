package _744_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSmallestLetterGreaterThanTarget(t *testing.T) {
	cases := []struct {
		name    string
		letters []byte
		target  byte
		res     byte
	}{
		{
			name:    "1",
			letters: []byte("cfj"),
			target:  byte('a'),
			res:     byte('c'),
		},
		{
			name:    "2",
			letters: []byte("cfj"),
			target:  byte('c'),
			res:     byte('f'),
		},
		{
			name:    "3",
			letters: []byte("xxyy"),
			target:  byte('z'),
			res:     byte('x'),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, nextGreatestLetter(c.letters, c.target))
		})
	}
}

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target {
		return letters[0]
	}

	for _, ltr := range letters {
		if ltr > target {
			return ltr
		}
	}

	return letters[0]
}
