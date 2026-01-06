package _2299_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/strong-password-checker-ii

func Test(t *testing.T) {
	cases := []struct {
		name string
		pwd  string
		res  bool
	}{
		{
			name: "1",
			pwd:  "IloveLe3tcode!",
			res:  true,
		},
		{
			name: "2",
			pwd:  "Me+You--IsMyDream",
			res:  false,
		},
		{
			name: "3",
			pwd:  "1aB!",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, strongPasswordCheckerII(c.pwd))
		})
	}
}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		uc, lc bool
		dig    bool
		spec   bool
		prev   rune
	)

	for _, r := range password {
		if r == prev {
			return false
		}
		prev = r

		if !uc && unicode.IsUpper(r) {
			uc = true
			continue
		}
		if !lc && unicode.IsLower(r) {
			lc = true
			continue
		}
		if !dig && unicode.IsDigit(r) {
			dig = true
			continue
		}

		if spec || unicode.IsLetter(r) {
			continue
		}

		switch r {
		case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+':
			spec = true
		}
	}

	return uc && lc && dig && spec
}
