package _504_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestBase7(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  string
	}{
		{
			name: "1",
			num:  100,
			res:  "202",
		},
		{
			name: "2",
			num:  -7,
			res:  "-10",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := convertToBase7(c.num)
			assert.Equal(t, c.res, res)
		})
	}
}

func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}
