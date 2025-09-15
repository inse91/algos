package _2409_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-days-spent-together

func Test(t *testing.T) {
	cases := []struct {
		name        string
		aliceArrive string
		aliceLeave  string
		bobArrive   string
		bobLeave    string
		res         int
	}{
		{
			name:        "1",
			aliceArrive: "08-15",
			aliceLeave:  "08-18",
			bobArrive:   "08-16",
			bobLeave:    "08-19",
			res:         3,
		},
		{
			name:        "2",
			aliceArrive: "10-01",
			aliceLeave:  "10-31",
			bobArrive:   "11-01",
			bobLeave:    "12-31",
			res:         0,
		},
		{
			name:        "11",
			aliceArrive: "08-15",
			aliceLeave:  "08-24",
			bobArrive:   "08-16",
			bobLeave:    "08-17",
			res:         2,
		},
		{
			name:        "12",
			aliceArrive: "07-16",
			aliceLeave:  "08-04",
			bobArrive:   "07-14",
			bobLeave:    "08-01",
			res:         17,
		},
		{
			name:        "25",
			aliceArrive: "10-20",
			aliceLeave:  "12-22",
			bobArrive:   "06-21",
			bobLeave:    "07-05",
			res:         0,
		},
		{
			name:        "49",
			aliceArrive: "01-20",
			aliceLeave:  "04-18",
			bobArrive:   "01-01",
			bobLeave:    "08-30",
			res:         89,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countDaysTogether(c.aliceArrive, c.aliceLeave, c.bobArrive, c.bobLeave))
		})
	}
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	const (
		format    = "01-02-2006"
		yearAffix = "-2001"
	)
	st1, _ := time.Parse(format, arriveAlice+yearAffix)
	st2, _ := time.Parse(format, arriveBob+yearAffix)
	if st2.After(st1) {
		st1 = st2
	}

	en1, _ := time.Parse(format, leaveAlice+yearAffix)
	en2, _ := time.Parse(format, leaveBob+yearAffix)
	if en2.Before(en1) {
		en1 = en2
	}

	hrs := en1.Sub(st1).Hours()
	if hrs < 0 {
		return 0
	}

	return int(hrs/24) + 1
}
