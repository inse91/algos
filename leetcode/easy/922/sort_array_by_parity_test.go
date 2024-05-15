package _922_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{4, 2, 5, 7},
			res:  []int{4, 5, 2, 7},
		},
		{
			name: "2",
			nums: []int{2, 3},
			res:  []int{2, 3},
		},
		{
			name: "28",
			nums: []int{648, 831, 560, 986, 192, 424, 997, 829, 897, 843},
			res:  []int{648, 831, 560, 997, 192, 897, 986, 829, 424, 843},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := sortArrayByParityII(c.nums)
			for i, n := range r {
				assert.Equal(t, c.res[i], n)
			}
		})
	}
}

func sortArrayByParityII(nums []int) []int {
	odd := make(map[int]struct{})
	even := make(map[int]struct{})
	for i, num := range nums {
		if i%2 == 0 && num%2 == 0 || i%2 == 1 && num%2 == 1 {
			continue
		}

		if i%2 == 0 {
			odd[i] = struct{}{}
		}
		if i%2 == 1 {
			even[i] = struct{}{}
		}

		if len(odd) != 0 && len(even) != 0 {
			var oddIdx int
			for k := range odd {
				oddIdx = k
				break
			}
			delete(odd, oddIdx)

			var evenIdx int
			for k := range even {
				evenIdx = k
				break
			}
			delete(even, evenIdx)
			nums[evenIdx], nums[oddIdx] = nums[oddIdx], nums[evenIdx]
		}
	}

	return nums
}
