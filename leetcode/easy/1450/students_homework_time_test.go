package _1450_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/

func Test(t *testing.T) {
	cases := []struct {
		name      string
		startTime []int
		endTime   []int
		queryTime int
		res       int
	}{
		{
			name:      "1",
			startTime: []int{1, 2, 3},
			endTime:   []int{3, 2, 7},
			queryTime: 4,
			res:       1,
		},
		{
			name:      "2",
			startTime: []int{4},
			endTime:   []int{4},
			queryTime: 4,
			res:       1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, busyStudent(c.startTime, c.endTime, c.queryTime))
		})
	}
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var res int

	for i := range startTime {
		if queryTime >= startTime[i] && endTime[i] >= queryTime {
			res++
		}
	}

	return res
}
