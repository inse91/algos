package _1576_test

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		s     string
		res   string
		noRes string
	}{
		{
			name:  "1",
			s:     "?zs",
			res:   "azs",
			noRes: "zzs",
		},
		{
			name:  "2",
			s:     "ubv?w",
			res:   "ubvaw",
			noRes: "ubvvw",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for range int(1e6) {
				s := modifyString(c.s)
				assert.NotEqual(t, c.noRes, s)
			}
		})
	}
}

var letters = [26]byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
	'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func getDifferByte(bts ...byte) byte {
	if len(bts) == 0 {
		return letters[0]
	}

	for {
		randIndex := rand.IntN(26)
		if letters[randIndex] == bts[0] {
			continue
		}

		if len(bts) == 1 || letters[randIndex] != bts[1] {
			return letters[randIndex]
		}
	}
}

func modifyString(s string) string {
	bts := []byte(s)

	bytesToIgnore := make([]byte, 0, 2)
	for i, b := range bts {
		if b != '?' {
			continue
		}

		if i != 0 {
			bytesToIgnore = append(bytesToIgnore, bts[i-1])
		}
		if i != len(bts)-1 {
			bytesToIgnore = append(bytesToIgnore, bts[i+1])
		}

		bts[i] = getDifferByte(bytesToIgnore...)
		bytesToIgnore = make([]byte, 0, 2)
	}

	return string(bts)
}
