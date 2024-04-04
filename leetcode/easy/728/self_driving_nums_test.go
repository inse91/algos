package _728_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSelfDividingNums(t *testing.T) {
	cases := []struct {
		name        string
		left, right int
		res         []int
	}{
		{
			name:  "1",
			left:  1,
			right: 22,
			res:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			name:  "2",
			left:  47,
			right: 85,
			res:   []int{48, 55, 66, 77},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := selfDividingNumbers(c.left, c.right)
			assert.Equal(t, c.res, res)
		})
	}
}

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0, (right-left)/3)

outer:
	for i := left; i <= right; i++ {
		if i < 10 {
			res = append(res, i)
			continue
		}
		digs := getDigits(i)
		for _, d := range digs {
			if d == 0 || i%d != 0 {
				continue outer
			}
		}
		res = append(res, i)
	}

	return res
}

func getDigits(in int) []int {
	s := strconv.Itoa(in)
	res := make([]int, 0, len(s))
	for _, r := range s {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		res = append(res, n)
	}

	return res
}
