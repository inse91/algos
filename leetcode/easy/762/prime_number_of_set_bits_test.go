package _762_test

import (
	"math"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		left  int
		right int
		res   int
	}{
		{
			name:  "1",
			left:  6,
			right: 10,
			res:   4,
		},
		{
			name:  "2",
			left:  10,
			right: 15,
			res:   5,
		},
		{
			name:  "3",
			left:  11,
			right: 17,
			res:   5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countPrimeSetBits(c.left, c.right))
		})
	}
}

func countPrimeSetBits(left int, right int) (res int) {
	for num := left; num <= right; num++ {
		onesCount := bits.OnesCount64(uint64(num))
		if isPrime(onesCount) {
			res++
		}
	}
	return res
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
