package _2651_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/calculate-delayed-arrival-time

func Test(t *testing.T) {
	cases := []struct {
		name  string
		time  int
		delay int
		res   int
	}{
		{
			name:  "1",
			time:  15,
			delay: 5,
			res:   20,
		},
		{
			name:  "2",
			time:  13,
			delay: 11,
			res:   0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findDelayedArrivalTime(c.time, c.delay))
		})
	}
}

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
