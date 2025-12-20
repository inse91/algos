package _1869_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longer-contiguous-segments-of-ones-than-zeros

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "1101",
			res:  true,
		},
		{
			name: "2",
			s:    "111000",
			res:  false,
		},
		{
			name: "3",
			s:    "110100010",
			res:  false,
		},
		{
			name: "13",
			s:    "0",
			res:  false,
		},
		{
			name: "14",
			s:    "1",
			res:  true,
		},
		{
			name: "127",
			s:    "0100011100",
			res:  false,
		},
		{
			name: "135",
			s:    "011000111",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, checkZeroOnes(c.s))
		})
	}
}

func checkZeroOnes(s string) bool {
	var (
		first = s[0]
		last  = s[len(s)-1]

		curLen = 1
		mx0    int
		mx1    int
	)

	if first == '0' {
		mx0 = 1
	} else {
		mx1 = 1
	}

	for i := 1; i < len(s); i++ {
		prev := s[i-1]
		cur := s[i]

		if cur == prev {
			curLen++
			continue
		}

		switch prev {
		case '0':
			mx0 = max(mx0, curLen)
		case '1':
			mx1 = max(mx1, curLen)
		}

		curLen = 1
	}

	if curLen < mx1 {
		return mx1 > mx0
	}

	switch last {
	case '0':
		mx0 = max(mx0, curLen)
	case '1':
		mx1 = max(mx1, curLen)
	}

	return mx1 > mx0
}
