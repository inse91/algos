package _1360_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-days-between-two-dates/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		d1   string
		d2   string
		res  int
	}{
		{
			name: "1",
			d1:   "2019-06-29",
			d2:   "2019-06-30",
			res:  1,
		},
		{
			name: "2",
			d1:   "2020-01-15",
			d2:   "2019-12-31",
			res:  15,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, daysBetweenDates(c.d1, c.d2))
		})
	}

}

func daysBetweenDates(date1 string, date2 string) int {
	d1, err := time.Parse(time.DateOnly, date1)
	if err != nil {
		panic(err)
	}

	d2, err := time.Parse(time.DateOnly, date2)
	if err != nil {
		panic(err)
	}

	sub := d2.Sub(d1)
	if sub < 0 {
		sub = -sub
	}

	days := sub.Hours() / 24

	return int(days)
}
