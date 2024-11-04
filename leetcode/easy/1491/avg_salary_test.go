package _1491_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/description/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		salary []int
		res    float64
	}{
		{
			name:   "1",
			salary: []int{4000, 3000, 1000, 2000},
			res:    2500.0000,
		},
		{
			name:   "2",
			salary: []int{1000, 3000, 2000},
			res:    2000.0000,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, average(c.salary))
		})
	}
}

func average(salary []int) float64 {
	mn, mx := int(10e6), 1000
	var sum int

	for _, s := range salary {
		sum += s
		mx = max(mx, s)
		mn = min(mn, s)
	}
	res := float64(sum-mx-mn) / float64(len(salary)-2)

	return res
}
