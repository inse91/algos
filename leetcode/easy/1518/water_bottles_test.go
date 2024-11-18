package _1518_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/water-bottles

func Test(t *testing.T) {
	cases := []struct {
		name    string
		bottles int
		num     int
		res     int
	}{
		{
			name:    "1",
			bottles: 9,
			num:     3,
			res:     13,
		},
		{
			name:    "2",
			bottles: 15,
			num:     4,
			res:     19,
		},
		{
			name:    "3",
			bottles: 18,
			num:     4,
			res:     23,
		},
		{
			name:    "101",
			bottles: 18,
			num:     10,
			res:     19,
		},
		{
			name:    "102",
			bottles: 19,
			num:     10,
			res:     21,
		},
		{
			name:    "103",
			bottles: 20,
			num:     10,
			res:     22,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numWaterBottles(c.bottles, c.num))
		})
	}
}

func numWaterBottles(numBottles int, numExchange int) int {
	var (
		res                 int
		emptyBottles        int
		uselessEmptyBottles int
	)

	for {
		res += numBottles
		if numBottles+uselessEmptyBottles < numExchange {
			break
		}

		emptyBottles = numBottles + uselessEmptyBottles
		uselessEmptyBottles = emptyBottles % numExchange
		numBottles = (emptyBottles - uselessEmptyBottles) / numExchange
	}

	return res
}

func numWaterBottlesBEST(numBottles int, numExchange int) int {
	return (numBottles*numExchange - 1) / (numExchange - 1)
}
