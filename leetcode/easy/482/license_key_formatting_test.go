package _482_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	cases := []struct {
		name string
		s    string
		k    int
		res  string
	}{
		{
			name: "1",
			s:    "5F3Z-2e-9-w",
			k:    4,
			res:  "5F3Z-2E9W",
		},
		{
			name: "2",
			s:    "2-5g-3-J",
			k:    2,
			res:  "2-5G-3J",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := licenseKeyFormatting(c.s, c.k)
			assert.Equal(t, c.res, res)
		})
	}
}

func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(s)
	ss := strings.Split(s, "-")
	s = strings.Join(ss, "")

	partsLen := len(s)/k + 1
	var firstPartLen int
	if len(s)%k == 0 {
		partsLen--
	} else {
		firstPartLen = len(s) % k
	}

	ss = make([]string, 0, partsLen)
	if firstPartLen != 0 {
		ss = append(ss, s[:firstPartLen])
		partsLen--
	}

	s = s[firstPartLen:]

	for i := 0; i < partsLen; i++ {
		end := (i + 1) * k
		st := end - k
		ss = append(ss, s[st:end])
	}

	return strings.Join(ss, "-")
}
