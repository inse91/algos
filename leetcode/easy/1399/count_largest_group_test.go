package _1399_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-largest-group/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    13,
			res:  4,
		},
		{
			name: "2",
			n:    2,
			res:  2,
		},
		{
			name: "22",
			n:    133,
			res:  1,
		},
		{
			name: "23",
			n:    134,
			res:  2,
		},
		{
			name: "24",
			n:    135,
			res:  1,
		},
		{
			name: "25",
			n:    1359,
			res:  1,
		},
		{
			name: "26",
			n:    2488,
			res:  1,
		},
		{
			name: "27",
			n:    23,
			res:  4,
		},
		{
			name: "28",
			n:    25,
			res:  6,
		},
		{
			name: "29",
			n:    37,
			res:  7,
		},
		{
			name: "30",
			n:    41,
			res:  2,
		},
		{
			name: "31",
			n:    10,
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countLargestGroup(c.n))
		})
	}
}

func countLargestGroup(n int) int {
	if n < 10 {
		return n
	}

	m := map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1, 9: 1}
	for i := 10; i <= n; i++ {
		sum := sumOfDigits(i)
		m[sum]++
	}
	var mx, res int
	for _, v := range m {
		if v < mx {
			continue
		}

		if v == mx {
			res++
			continue
		}

		res = 1
		mx = v
	}

	return res
}

func sumOfDigits(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}

func TestSumOfDigits(t *testing.T) {
	assert.Equal(t, 6, sumOfDigits(123))
	assert.Equal(t, 16, sumOfDigits(448))
}
