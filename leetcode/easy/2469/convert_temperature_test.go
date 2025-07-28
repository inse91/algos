package _2469_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-the-temperature

func Test(t *testing.T) {
	cases := []struct {
		name      string
		c float64
		res       []float64
	}{
		{
			name: "1",
			c: 36.50,
			res: []float64{309.65000,97.70000},
		},
		{
			name: "2",
			c: 122.11,
			res: []float64{395.26000,251.79800},
		},

	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, convertTemperature(c.c))
		})
	}
}

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius+273.15, celsius*1.8+32}
}
