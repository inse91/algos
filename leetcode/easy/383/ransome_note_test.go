package _383_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRansomNote(t *testing.T) {
	cases := []struct {
		name string
		in1  string
		in2  string
		res  bool
	}{
		{
			name: "1",
			in1:  "a",
			in2:  "b",
			res:  false,
		},
		{
			name: "2",
			in1:  "aa",
			in2:  "ab",
			res:  false,
		},
		{
			name: "3",
			in1:  "aa",
			in2:  "aab",
			res:  true,
		},
		{
			name: "4",
			in1:  "a",
			in2:  "a",
			res:  true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := canConstruct(c.in1, c.in2)
			assert.Equal(t, c.res, res)
		})
	}
}

func canConstruct(ransomNote string, magazine string) bool {
	store := make(map[byte]int, 16)
	for _, r := range magazine {
		if _, ok := store[byte(r)]; !ok {
			store[byte(r)] = 1
			continue
		}
		store[byte(r)]++
	}
	for _, r := range ransomNote {
		if q, ok := store[byte(r)]; !ok || q == 0 {
			return false
		}
		store[byte(r)]--
	}

	return true
}
