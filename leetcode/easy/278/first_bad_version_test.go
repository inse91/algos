package _278_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  int
	}{
		{
			name: "1",
			in:   5,
			res:  4,
		},
		{
			name: "2",
			in:   1,
			res:  1,
		},
		{
			name: "3",
			in:   2,
			res:  1,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := firstBadVersion(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	v := n
	for isBadVersion(v) && v > 1 {
		v--
	}

	if v < n {
		return v + 1
	}
	return n
}

func isBadVersion(version int) bool {
	switch version {
	case 1, 2, 4, 5:
		return true
	default:
		return false
	}
}
