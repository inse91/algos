package _228_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		res  []string
	}{
		{
			name: "1",
			in:   []int{0, 1, 2, 4, 5, 7},
			res:  []string{"0->2", "4->5", "7"},
		},
		{
			name: "2",
			in:   []int{0, 2, 3, 4, 6, 8, 9},
			res:  []string{"0", "2->4", "6", "8->9"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := summaryRanges(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}

func summaryRanges(nums []int) []string {
	if nums == nil {
		return nil
	}
	if len(nums) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	st := nums[0]
	fn := st
	for i := 1; i < len(nums); i++ {
		if nums[i] == fn+1 {
			fn++
			continue
		}
		if st == fn {
			res = append(res, strconv.Itoa(st))
			st = nums[i]
			fn = st
			continue
		}
		res = append(res, strconv.Itoa(st)+"->"+strconv.Itoa(fn))
		st = nums[i]
		fn = st
	}
	if st == fn {
		res = append(res, strconv.Itoa(st))
	} else {
		res = append(res, strconv.Itoa(st)+"->"+strconv.Itoa(fn))
	}
	return res
}
