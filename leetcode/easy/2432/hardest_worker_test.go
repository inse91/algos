package _2432_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		logs [][]int
		res  int
	}{
		{
			name: "1",
			n:    10,
			logs: [][]int{
				{0, 3},
				{2, 5},
				{0, 9},
				{1, 15},
			},
			res: 1,
		},
		{
			name: "2",
			n:    26,
			logs: [][]int{
				{1, 1},
				{3, 7},
				{2, 12},
				{7, 17},
			},
			res: 3,
		},
		{
			name: "3",
			n:    2,
			logs: [][]int{
				{0, 10},
				{1, 20},
			},
			res: 0,
		},
		{
			name: "70",
			n:    70,
			logs: [][]int{
				{36, 3},
				{1, 5},
				{12, 8},
				{25, 9},
				{53, 11},
				{29, 12},
				{52, 14},
			},
			res: 12,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, hardestWorker(c.n, c.logs))
		})
	}
}

func hardestWorker(_ int, logs [][]int) int {
	maxTime := logs[0][1]
	maxUserId := logs[0][0]

	for taskId, log := range logs {
		if taskId == 0 {
			continue
		}

		currentUserId := log[0]
		startTime := logs[taskId-1][1]
		endTime := log[1]

		taskTime := endTime - startTime
		if taskTime < maxTime {
			continue
		}

		if taskTime == maxTime && currentUserId < maxUserId {
			maxUserId = currentUserId
		}

		if taskTime > maxTime {
			maxTime = taskTime
			maxUserId = currentUserId
		}
	}

	return maxUserId
}
