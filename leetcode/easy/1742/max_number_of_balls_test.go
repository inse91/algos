package _1742_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-number-of-balls-in-a-box/description/

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
			high: 10,
			res:  2,
		},
		{
			name: "2",
			low:  5,
			high: 15,
			res:  2,
		},
		{
			name: "3",
			low:  19,
			high: 28,
			res:  2,
		},
		{
			name: "64",
			low:  52,
			high: 61,
			res:  2,
		},
		{
			name: "69",
			low:  133,
			high: 299,
			res:  17,
		},
		{
			name: "71",
			low:  1025,
			high: 3030,
			res:  150,
		},
		{
			name: "83",
			low:  1238,
			high: 5317,
			res:  316,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countBalls(c.low, c.high))
		})
	}
}

func countBalls(lowLimit int, highLimit int) int {
	var res int
	bag := make(map[int]int)
	n := highLimit - lowLimit + 1
	for i := 1; i <= n; i++ {
		sm := getSumOfDigits(lowLimit)
		lowLimit++

		bag[sm]++
		res = max(res, bag[sm])
	}

	return res
}

func getSumOfDigits(n int) int {
	var acc int
	for n > 0 {
		d := n / 10

		acc += n - d*10
		n /= 10
	}

	return acc
}
