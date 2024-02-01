package _495_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTeemoAttacking(t *testing.T) {
	cases := []struct {
		name       string
		timeSeries []int
		duration   int
		res        int
	}{
		{
			name:       "1",
			timeSeries: []int{1, 4},
			duration:   2,
			res:        4,
		},
		{
			name:       "2",
			timeSeries: []int{1, 2},
			duration:   2,
			res:        3,
		},
		{
			name:       "3",
			timeSeries: []int{1, 2},
			duration:   1,
			res:        2,
		},
		{
			name:       "4",
			timeSeries: []int{1, 2},
			duration:   3,
			res:        4,
		},
		{
			name:       "5",
			timeSeries: []int{1, 2},
			duration:   4,
			res:        5,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findPoisonedDuration(c.timeSeries, c.duration)
			assert.Equal(t, c.res, res)
		})
	}
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	var res int
	for i := 1; i < len(timeSeries); i++ {
		gap := timeSeries[i] - timeSeries[i-1]
		res += min(duration, gap)
	}

	return res + duration
}
