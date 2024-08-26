package _1281_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    234,
			res:  15,
		},
		{
			name: "2",
			n:    4421,
			res:  21,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, subtractProductAndSum(c.n))
		})
	}
}

func subtractProductAndSum(n int) int {
	var prod int = 1
	var sum int

	var c int
	var n1 int
	for n > 0 {
		n1 = n / 10
		c = n - n1*10

		sum += c
		prod *= c

		n = n1
	}

	return prod - sum
}
