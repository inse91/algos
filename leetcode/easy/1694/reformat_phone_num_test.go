package _1694_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reformat-phone-number/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		phone string
		res   string
	}{
		{
			name:  "1",
			phone: "1-23-45 6",
			res:   "123-456",
		},
		{
			name:  "2",
			phone: "123 4-567",
			res:   "123-45-67",
		},
		{
			name:  "3",
			phone: "123 4-5678",
			res:   "123-456-78",
		},
		{
			name:  "13",
			phone: "1234",
			res:   "12-34",
		},
		{
			name:  "17",
			phone: "1-23",
			res:   "123",
		},
		{
			name:  "19",
			phone: "1 2",
			res:   "12",
		},
		{
			name:  "23",
			phone: "12345",
			res:   "123-45",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, reformatNumber(c.phone))
		})
	}
}

func reformatNumber(number string) string {
	bts := make([]byte, 0, len(number))
	for _, r := range number {
		if !unicode.IsDigit(r) {
			continue
		}
		bts = append(bts, byte(r))
	}

	number = string(bts)

	ln := len(number)
	switch ln % 3 {
	case 0:
		return splitInto(number, 3)
	case 1:
		start := number[:ln-4]
		end := number[ln-4:]
		start = splitInto(start, 3)
		end = splitInto(end, 2)

		if start == "" {
			return end
		}

		return start + "-" + end
	case 2:
		start := number[:ln-2]
		end := number[ln-2:]
		start = splitInto(start, 3)

		if start == "" {
			return end
		}

		return start + "-" + end
	}

	return ""
}

func splitInto(s string, n int) string {
	if len(s) == 0 {
		return ""
	}

	var sb strings.Builder
	var c int
	sb.Grow(len(s) + len(s)/n - 1)
	for _, r := range s {
		if c == n {
			sb.WriteByte('-')
			c = 0
		}

		sb.WriteByte(byte(r))
		c++
	}

	return sb.String()
}
