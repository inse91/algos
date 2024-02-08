package _541__test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"strings"
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		name string
		s    string
		k    int
		res  string
	}{
		{
			name: "1",
			s:    "abcdefg",
			k:    2,
			res:  "bacdfeg",
		},
		{
			name: "2",
			s:    "abcd",
			k:    2,
			res:  "bacd",
		},
		{
			name: "3",
			s:    "abcdabcd",
			k:    2,
			res:  "bacdbacd",
		},
		{
			name: "4",
			s:    "abcdabcdab",
			k:    2,
			res:  "bacdbacdba",
		},
		{
			name: "5",
			s:    "abcdabcdabc",
			k:    2,
			res:  "bacdbacdbac",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := reverseStr(c.s, c.k)
			assert.Equal(t, c.res, res)
		})
	}
}

func reverseStr(s string, k int) string {
	parts := splitSrt(s, 2*k)

	// all except last one
	for i, pc := range parts[:len(parts)-1] {
		firstPart := pc[:k]
		bts := []byte(firstPart)
		slices.Reverse(bts)

		parts[i] = string(bts) + pc[k:]
	}

	// last one
	lastPc := parts[len(parts)-1]
	if len(lastPc) >= k {
		firstPart := lastPc[:k]
		bts := []byte(firstPart)
		slices.Reverse(bts)

		parts[len(parts)-1] = string(bts) + lastPc[k:]
	} else {
		bts := []byte(lastPc)
		slices.Reverse(bts)

		parts[len(parts)-1] = string(bts)
	}

	return strings.Join(parts, "")
}

func splitSrt(s string, pcLen int) []string {
	res := make([]string, 0, len(s)/pcLen)
	for i := range s {
		if (i+1)%pcLen == 0 {
			end := i + 1
			start := end - pcLen
			res = append(res, s[start:end])
		}
	}

	if len(s)%pcLen != 0 {
		end := len(s)
		start := end - len(s)%pcLen
		res = append(res, s[start:end])
	}

	return res
}
