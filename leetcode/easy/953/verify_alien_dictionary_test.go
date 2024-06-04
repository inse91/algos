package _953_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		order string
		res   bool
	}{
		{
			name:  "1",
			words: []string{"hello", "leetcode"},
			order: "hlabcdefgijkmnopqrstuvwxyz",
			res:   true,
		},
		{
			name:  "2",
			words: []string{"word", "world", "row"},
			order: "worldabcefghijkmnpqstuvxyz",
			res:   false,
		},
		{
			name:  "3",
			words: []string{"apple", "app"},
			order: "abcdefghijklmnopqrstuvwxyz",
			res:   false,
		},
		{
			name:  "28",
			words: []string{"kuvp", "q"},
			order: "ngxlkthsjuoqcpavbfdermiywz",
			res:   true,
		},
		{
			name:  "29",
			words: []string{"app", "apple"},
			order: "abcdefghijklmnopqrstuvwxyz",
			res:   true,
		},
		{
			name:  "113",
			words: []string{"ubg", "kwh"},
			order: "qcipyamwvdjtesbghlorufnkzx",
			res:   true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isAlienSorted(c.words, c.order))
		})
	}
}

func isAlienSorted(words []string, order string) bool {
	if len(words) <= 1 {
		return true
	}

	weights := make(map[rune]int8, 26)
	for i, r := range order {
		weights[r] = int8(i)
	}

	for i := 1; i < len(words); i++ {
		cur := words[i]
		prev := words[i-1]

		for idx, r := range cur {
			if idx >= len(prev) {
				break
			}

			curW, ok := weights[r]
			if !ok {
				panic("invalid char: " + string(r))
			}
			prevW, ok := weights[rune(prev[idx])]
			if !ok {
				panic("invalid char: " + string(prev[idx]))
			}

			if curW > prevW {
				break
			}

			if curW == prevW {
				if idx != len(cur)-1 {
					continue
				}

				if len(prev) > len(cur) {
					return false
				}

			}

			if curW < prevW {
				return false
			}
		}
	}

	return true
}
