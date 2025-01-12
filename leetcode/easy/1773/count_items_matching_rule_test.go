package _1773_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/largest-substring-between-two-equal-characters/description/

func Test(t *testing.T) {
	cases := []struct {
		name      string
		items     [][]string
		ruleKey   string
		ruleValue string
		res       int
	}{
		{
			name: "1",
			items: [][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "lenovo"},
				{"phone", "gold", "iphone"},
			},
			ruleKey:   "color",
			ruleValue: "silver",
			res:       1,
		},
		{
			name: "2",
			items: [][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "lenovo"},
				{"phone", "gold", "iphone"},
			},
			ruleKey:   "type",
			ruleValue: "phone",
			res:       2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countMatches(c.items, c.ruleKey, c.ruleValue))
		})
	}
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var keyIndex int
	switch ruleKey {
	case "color":
		keyIndex = 1
	case "name":
		keyIndex = 2
	}

	var c int
	for _, dvc := range items {
		if dvc[keyIndex] == ruleValue {
			c++
		}
	}

	return c
}
