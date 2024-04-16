package _693_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinNumWithAltBits(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  bool
	}{
		{
			name: "1",
			n:    5,
			res:  true,
		},
		{
			name: "2",
			n:    7,
			res:  false,
		},
		{
			name: "3",
			n:    11,
			res:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := hasAlternatingBits(c.n)
			assert.Equal(t, c.res, res)
		})
	}
}

func hasAlternatingBits(n int) bool {
	//s := strconv.FormatInt(int64(n), 2)
	//for i := range s {
	//	if i == 0 {
	//		continue
	//	}
	//	if s[i] == s[i-1] {
	//		return false
	//	}
	//}
	//return true

	x := -1

	for n > 0 {
		if x == -1 {
			x = n % 2
			n /= 2
			continue
		}
		if x == n%2 {
			return false
		}
		x = n % 2
		n /= 2

	}

	return true
}
