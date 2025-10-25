package _2843_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-symmetric-integers

func Test(t *testing.T) {
	cases := []struct {
		name string
		low  int
		high int
		res  int
	}{
		{
			name: "1",
			low:  1,
			high: 100,
			res:  9,
		},
		{
			name: "2",
			low:  1200,
			high: 1230,
			res:  4,
		},
		{
			name: "27",
			low:  44,
			high: 8675,
			res:  554,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countSymmetricIntegers(c.low, c.high))
		})
	}
}

func countSymmetricIntegers(low int, high int) int {
	var res int
	for _, num := range [...]int{11, 22, 33, 44, 55, 66, 77, 88, 99} {
		if low <= num && num <= high {
			res++
		}
	}

	if high <= 1000 {
		return res
	}
	if high == 10_000 {
		high--
	}

	low = max(low, 1001)
	for num := low; num <= high; num++ {
		p1 := num / 100
		p2 := num - p1*100

		sum1 := digSum(p1)
		sum2 := digSum(p2)
		if sum1 != sum2 {
			continue
		}

		res++
	}

	return res
}

func digSum(num int) int {
	first := num / 10
	sec := num - first*10

	return first + sec
}
