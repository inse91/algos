package _125

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"slices"
	"strings"
	"testing"
	"unicode"
)

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		name   string
		in     string
		expRes bool
	}{
		{
			name:   "empty_string",
			in:     " ",
			expRes: true,
		},
		{
			name:   "valid",
			in:     "A man, a plan, a canal: Panama",
			expRes: true,
		},
		{
			name:   "!valid",
			in:     "race a car",
			expRes: false,
		},
		{
			name:   "!valid",
			in:     "0P",
			expRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isPalindrome(tc.in)
			assert.Equal(t, tc.expRes, res)
		})
	}
}

func isPalindrome(s string) bool {
	sb := strings.Builder{}
	sb.Grow(len(s))

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			continue
		}
		if unicode.IsUpper(r) {
			sb.WriteString(strings.ToLower(string(r)))
			continue
		}
		sb.WriteRune(r)
	}

	s = sb.String()

	bts := []byte(s)
	slices.Reverse(bts)

	return string(bts) == sb.String()
}

func TestIsLetter(t *testing.T) {
	for _, r := range "this is a world" {
		fmt.Println(string(r), unicode.IsLetter(r))

	}

	s := []byte("123")
	slices.Reverse(s)

	fmt.Println(string(s))

}
