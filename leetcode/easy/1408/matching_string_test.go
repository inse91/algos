package _1408_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/string-matching-in-an-array

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   []string
	}{
		{
			name:  "1",
			words: []string{"mass", "as", "hero", "superhero"},
			res:   []string{"as", "hero"},
		},
		{
			name:  "2",
			words: []string{"leetcode", "et", "code"},
			res:   []string{"et", "code"},
		},
		{
			name:  "3",
			words: []string{"blue", "green", "bu"},
			res:   []string{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, stringMatching(c.words))
		})
	}
}

func stringMatching(words []string) []string {
	set := make(map[string]struct{}, len(words))
	for _, w := range words {
		set[w] = struct{}{}
	}

	var res []string
	banned := map[string]struct{}{}
	for _, w := range words {
		for k := range set {
			if _, ok := banned[k]; ok {
				continue
			}
			if k == w {
				continue
			}

			if strings.Contains(w, k) {
				res = append(res, k)
				banned[k] = struct{}{}
			}
		}
	}

	return res
}
