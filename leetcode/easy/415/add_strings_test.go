package _415_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"strings"
	"testing"
)

func TestAddStrings(t *testing.T) {
	cases := []struct {
		name string
		in1  string
		in2  string
		res  string
	}{
		{
			name: "1",
			in1:  "11",
			in2:  "123",
			res:  "134",
		},
		{
			name: "2",
			in1:  "456",
			in2:  "77",
			res:  "533",
		},
		{
			name: "3",
			in1:  "0",
			in2:  "0",
			res:  "0",
		},
		{
			name: "4",
			in1:  "9",
			in2:  "1",
			res:  "10",
		},
		{
			name: "5",
			in1:  "6913259244",
			in2:  "71103343",
			res:  "6984362587",
		},
		{
			name: "6",
			in1:  "3876620623801494171",
			in2:  "6529364523802684779",
			res:  "10405985147604178950",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := addStrings(c.in1, c.in2)
			assert.Equal(t, c.res, res)
		})
	}
}

func addStrings(num1 string, num2 string) string {
	maxLen := max(len(num1), len(num2))
	diff := maxLen - min(len(num1), len(num2))
	prefix := strings.Repeat("0", diff)
	if len(num1) < maxLen {
		num1 = prefix + num1
	} else {
		num2 = prefix + num2
	}

	const byteGap byte = 48
	sb := strings.Builder{}
	var r uint8

	for i := len(num1) - 1; i >= 0; i-- {
		sum := num1[i] + num2[i] + r - 2*byteGap
		if sum > 9 {
			sum -= 10
			r = 1
		} else {
			r = 0
		}
		sb.WriteByte(sum + byteGap)
	}
	if r != 0 {
		sb.WriteByte(r + byteGap)
	}

	bts := []byte(sb.String())
	slices.Reverse(bts)
	s := string(bts)
	return s
}
