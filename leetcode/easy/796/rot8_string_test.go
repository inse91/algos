package _796_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateString(t *testing.T) {
	cases := []struct {
		name string
		s    string
		goal string
		res  bool
	}{
		{
			name: "1",
			s:    "abcde",
			goal: "cdeab",
			res:  true,
		},
		{
			name: "2",
			s:    "abcde",
			goal: "abced",
			res:  false,
		},
		{
			name: "29",
			s:    "clrwmpkwru",
			goal: "wmpkwruclr",
			res:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, rotateString(c.s, c.goal))
		})
	}
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}

	for i := range s {
		partS := s[i:]
		partGoal := goal[:len(goal)-i]
		if partS != partGoal {
			continue
		}
		partS = s[:i]
		partGoal = goal[len(goal)-i:]
		if partS == partGoal {
			return true
		}

	}

	return false
}
