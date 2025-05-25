package _3174_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/clear-digits

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "abc",
			res:  "abc",
		},
		{
			name: "2",
			s:    "cb34",
			res:  "",
		},
		{
			name: "3",
			s:    "cb34a5c",
			res:  "c",
		},
		{
			name: "4",
			s:    "dbf34a5",
			res:  "d",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, clearDigits(c.s))
		})
	}
}

func clearDigits(s string) string {
	var digsIdxs []int
	idxs := make([]bool, len(s)) // если true то исключаем из итоговой строки
	for i, r := range s {
		if !unicode.IsDigit(r) {
			continue
		}

		digsIdxs = append(digsIdxs, i)
		idxs[i] = true
	}

	if len(digsIdxs) == 0 {
		return s
	}

	for _, idx := range digsIdxs {
		for prev := idx - 1; prev >= 0; prev-- {
			if idxs[prev] {
				continue
			}

			idxs[prev] = true
			break
		}
	}

	var sb strings.Builder
	for i, r := range s {
		if !idxs[i] {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
