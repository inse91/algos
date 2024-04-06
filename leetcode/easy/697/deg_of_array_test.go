package _697_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDegOfArray(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 2, 3, 1},
			res:  2,
		},
		{
			name: "2",
			nums: []int{1, 2, 2, 3, 1, 4, 2},
			res:  6,
		},
		{
			name: "3",
			nums: []int{1, 2, 1},
			res:  3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := findShortestSubArray(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func findShortestSubArray(nums []int) int {
	indexes := make(map[int][]int)
	maxes := make(map[int]int)

	for i, num := range nums {
		v, ok := indexes[num]
		if !ok {
			indexes[num] = []int{i}
			maxes[num] = 1
			continue
		}

		indexes[num] = append(v, i)
		maxes[num]++
	}

	var mx int
	var mxx []int
	for k, v := range maxes {
		if v < mx {
			continue
		}
		if v == mx {
			mxx = append(mxx, k)
			continue
		}
		// v > mx
		mx = v
		mxx = []int{k}
	}

	res := len(nums) + 1
	for _, key := range mxx {
		ixes := indexes[key]
		if len(ixes) == 0 {
			continue
		}
		left := ixes[0]
		right := ixes[len(ixes)-1]

		l := right - left + 1
		res = min(res, l)
	}

	return res
}
