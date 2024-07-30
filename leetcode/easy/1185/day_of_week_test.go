package _1185_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// https://leetcode.com/problems/day-of-the-week/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		day   int
		month int
		year  int
		res   string
	}{
		{
			name:  "sat",
			day:   31,
			month: 8,
			year:  2019,
			res:   "Saturday",
		},
		{
			name:  "sun1",
			day:   18,
			month: 7,
			year:  1999,
			res:   "Sunday",
		},
		{
			name:  "sun2",
			day:   15,
			month: 8,
			year:  1993,
			res:   "Sunday",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, dayOfTheWeek(c.day, c.month, c.year))
		})
	}
}

func dayOfTheWeek(day int, month int, year int) string {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return date.Weekday().String()
}
