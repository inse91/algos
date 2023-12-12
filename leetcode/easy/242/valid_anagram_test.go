package _242_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mS := map[byte]int{}
	mT := map[byte]int{}

	for i := range s {
		if v, ok := mS[s[i]]; ok {
			mS[s[i]] = v + 1
		} else {
			mS[s[i]] = 1
		}
		if v, ok := mT[t[i]]; ok {
			mT[t[i]] = v + 1
		} else {
			mT[t[i]] = 1
		}
	}

	for k, v := range mS {
		val, ok := mT[k]
		if !ok || v != val {
			return false
		}
	}
	return true
}

func TestValidAnagram(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		t    string
		res  bool
	}{
		{
			name: "1",
			s:    "anagram",
			t:    "nagaram",
			res:  true,
		},
		{
			name: "2",
			s:    "rat",
			t:    "car",
			res:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isAnagram(tc.s, tc.t)
			assert.Equal(t, tc.res, res)
		})
	}
}
