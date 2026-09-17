package _2960_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-tested-devices-after-test-operations/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		percs []int
		res   int
	}{
		{
			name:  "1",
			percs: []int{1, 1, 2, 1, 3},
			res:   3,
		},
		{
			name:  "2",
			percs: []int{0, 1, 2},
			res:   2,
		},
		{
			name:  "54",
			percs: []int{2, 1},
			res:   1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countTestedDevices(append([]int{}, c.percs...)))
			assert.Equal(t, c.res, countTestedDevicesV2(append([]int{}, c.percs...)))
		})
	}
}

func countTestedDevices(pers []int) int {
	var (
		res  int
		rest []int
	)
	for i, perc := range pers {
		if perc <= 0 {
			continue
		}

		res += 1
		rest = pers[i+1:]
		for j := range rest {
			rest[j] -= 1
		}
	}

	return res
}

func countTestedDevicesV2(pers []int) int {
	var res int
	for _, perc := range pers {
		perc -= res
		if perc > 0 {
			res += 1
		}
	}

	return res
}
