package _551_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAttendanceRecord(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "PPALLP",
			res:  true,
		},
		{
			name: "2",
			s:    "PPALLL",
			res:  false,
		},
		{
			name: "3",
			s:    "PPALLPA",
			res:  false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := checkRecord(c.s)
			assert.Equal(t, c.res, res)
		})
	}
}

func checkRecord(s string) bool {
	const (
		abs  = "A"
		late = "LLL"
	)

	return strings.Count(s, abs) < 2 && !strings.Contains(s, late)
}
