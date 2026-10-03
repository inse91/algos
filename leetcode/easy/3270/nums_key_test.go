package _3270_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-key-of-the-numbers

func Test(t *testing.T) {
	cases := []struct {
		name             string
		num1, num2, num3 int
		res              int
	}{
		{
			name: "1",
			num1: 1,
			num2: 10,
			num3: 1000,
			res:  0,
		},
		{
			name: "2",
			num1: 987,
			num2: 879,
			num3: 798,
			res:  777,
		},
		{
			name: "3",
			num1: 1,
			num2: 2,
			num3: 3,
			res:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, generateKey(c.num1, c.num2, c.num3))
		})
	}
}

func generateKey(num1 int, num2 int, num3 int) int {
	var res int
	for i := 1; i <= 4; i++ {
		inc := min(take(num1, i), take(num2, i), take(num3, i))
		if inc == 0 {
			continue
		}

		for range 4 - i {
			inc *= 10
		}

		res += inc

	}

	return res
}

func take(num int, k int) int {
	switch k {
	case 1:
		num -= num / 10000 * 10000
		return num / 1000
	case 2:
		num -= num / 1000 * 1000
		return num / 100
	case 3:
		num -= num / 100 * 100
		return num / 10
	case 4:
		return num - num/10*10
	default:
		panic("invalid k")
	}
}

func TestTake(t *testing.T) {
	assert.Equal(t, 3, take(3856, 1))
	assert.Equal(t, 8, take(3856, 2))
	assert.Equal(t, 5, take(3856, 3))
	assert.Equal(t, 6, take(3856, 4))

	assert.Panics(t, func() { take(3856, 0) })
	assert.Panics(t, func() { take(3856, -13) })
	assert.Panics(t, func() { take(3856, 5) })
	assert.Panics(t, func() { take(3856, 24) })
}
