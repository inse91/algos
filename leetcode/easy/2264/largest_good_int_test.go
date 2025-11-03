package _2264_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/largest-3-same-digit-number-in-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  string
		res  string
	}{
		{
			name: "1",
			num:  "6777133339",
			res:  "777",
		},
		{
			name: "2",
			num:  "2300019",
			res:  "000",
		},
		{
			name: "3",
			num:  "42352338",
			res:  "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, largestGoodInteger(c.num))
		})
	}
}

func largestGoodInteger(num string) string {
	if len(num) <= 2 {
		return ""
	}

	var (
		found bool
		mx    byte
	)
	for i := 2; i < len(num); i++ {
		var (
			cur   = num[i]
			prev  = num[i-1]
			pprev = num[i-2]
		)

		if cur != prev || cur != pprev {
			continue
		}

		found = true
		mx = max(mx, cur)
		if mx == '9' {
			break
		}
	}

	if !found {
		return ""
	}

	return string(bytes.Repeat([]byte{mx}, 3))
}
