package _1859_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sorting-the-sentence

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "is2 sentence4 This1 a3",
			res:  "This is a sentence",
		},
		{
			name: "2",
			s:    "Myself2 Me1 I4 and3",
			res:  "Me Myself and I",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sortSentence(c.s))
		})
	}
}

func sortSentence(s string) string {
	pts := strings.Split(s, " ")
	wantPts := make([]string, len(pts))

	for _, pt := range pts {
		idx := int(pt[len(pt)-1] - '0' - 1)
		wantPts[idx] = pt[:len(pt)-1]
	}

	return strings.Join(wantPts, " ")
}
