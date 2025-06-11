package _2806_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/account-balance-after-rounded-purchase

func Test(t *testing.T) {
	cases := []struct {
		name string
		amt  int
		res  int
	}{
		{
			name: "1",
			amt:  9,
			res:  90,
		},
		{
			name: "2",
			amt:  15,
			res:  80,
		},
		{
			name: "3",
			amt:  10,
			res:  90,
		},
		{
			name: "14",
			amt:  55,
			res:  40,
		},
		{
			name: "61",
			amt:  11,
			res:  90,
		},
		{
			name: "62",
			amt:  21,
			res:  80,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, accountBalanceAfterPurchase(c.amt))
		})
	}
}

func accountBalanceAfterPurchase(purchaseAmount int) int {
	if purchaseAmount%10 != 0 {
		rnd := math.Round(float64(purchaseAmount) / 10)
		purchaseAmount = int(rnd) * 10
	}

	return 100 - purchaseAmount
}
