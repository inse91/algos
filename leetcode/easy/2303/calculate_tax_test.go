package _2303_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/calculate-amount-paid-in-taxes

func Test(t *testing.T) {
	cases := []struct {
		name     string
		brackets [][]int
		income   int
		res      float64
	}{
		{
			name: "1",
			brackets: [][]int{
				{3, 50},
				{7, 10},
				{12, 25},
			},
			income: 10,
			res:    2.65,
		},
		{
			name: "2",
			brackets: [][]int{
				{1, 0},
				{4, 25},
				{5, 50},
			},
			income: 2,
			res:    0.25,
		},
		{
			name: "3",
			brackets: [][]int{
				{2, 50},
			},
			income: 0,
			res:    0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, calculateTax(c.brackets, c.income))
		})
	}
}

func calculateTax(brackets [][]int, income int) float64 {
	if income == 0 {
		return 0
	}

	if len(brackets) == 0 {
		return float64(income)
	}

	var res float64
	for i, br := range brackets {
		var prev int
		if i != 0 {
			prev = brackets[i-1][0]
		}

		taxed := min(br[0], income) - prev
		res += float64(taxed) * float64(br[1]) / 100
		if income < br[0] {
			break
		}
	}

	return res
}
