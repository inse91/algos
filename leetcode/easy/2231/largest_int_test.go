package _2231_test

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/largest-number-after-digit-swaps-by-parity

func Test(t *testing.T) {
	cases := []struct {
		name string
		num  int
		res  int
	}{
		{
			name: "1",
			num:  1234,
			res:  3412,
		},
		{
			name: "2",
			num:  65875,
			res:  87655,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, largestInteger(c.num))
		})
	}
}

func largestInteger(num int) int {
	var digs [10]int
	order := make([]bool, 0, 9)

	for num > 0 {
		dig := num % 10
		order = append(order, dig%2 == 0)
		digs[dig]++
		num /= 10
	}

	var sb strings.Builder
	sb.Grow(len(order))

	for _, isEven := range slices.Backward(order) {
		if isEven {
			if digs[8] != 0 {
				digs[8]--
				sb.WriteRune('8')
				continue
			}
			if digs[6] != 0 {
				digs[6]--
				sb.WriteRune('6')
				continue
			}
			if digs[4] != 0 {
				digs[4]--
				sb.WriteRune('4')
				continue
			}
			if digs[2] != 0 {
				digs[2]--
				sb.WriteRune('2')
				continue
			}
			if digs[0] != 0 {
				digs[0]--
				sb.WriteRune('0')
				continue
			}

			panic("no even digits left")
		}

		if digs[9] != 0 {
			digs[9]--
			sb.WriteRune('9')
			continue
		}
		if digs[7] != 0 {
			digs[7]--
			sb.WriteRune('7')
			continue
		}
		if digs[5] != 0 {
			digs[5]--
			sb.WriteRune('5')
			continue
		}
		if digs[3] != 0 {
			digs[3]--
			sb.WriteRune('3')
			continue
		}
		if digs[1] != 0 {
			digs[1]--
			sb.WriteRune('1')
			continue
		}

		panic("no odd digits left")
	}

	res, err := strconv.Atoi(sb.String())
	if err != nil {
		panic(err)
	}

	return res
}
