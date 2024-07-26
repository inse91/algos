package _1154_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// https://leetcode.com/problems/day-of-the-year/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		date string
		res  int
	}{
		{
			name: "1",
			date: "2019-01-09",
			res:  9,
		},
		{
			name: "2",
			date: "2019-02-10",
			res:  41,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, dayOfYear(c.date))
		})
	}
}

func dayOfYear(date string) int {
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		panic(err)
	}

	dayOneOfYear := time.Date(d.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	hrs := d.Sub(dayOneOfYear).Hours()

	return int(hrs)/24 + 1
}
