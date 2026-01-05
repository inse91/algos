package _2446_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/determine-if-two-events-have-conflict

func Test(t *testing.T) {
	cases := []struct {
		name   string
		e1, e2 []string
		res    bool
	}{
		{
			name: "1",
			e1:   []string{"01:15", "02:00"},
			e2:   []string{"02:00", "03:00"},
			res:  true,
		},
		{
			name: "2",
			e1:   []string{"01:00", "02:00"},
			e2:   []string{"01:20", "03:00"},
			res:  true,
		},
		{
			name: "3",
			e1:   []string{"10:00", "11:00"},
			e2:   []string{"14:00", "15:00"},
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, haveConflict(c.e1, c.e2))
		})
	}
}

func haveConflict(event1 []string, event2 []string) bool {
	if len(event1) < 2 || len(event2) < 2 {
		panic("invalid events")
	}

	beg1 := getTime(event1[0])
	end1 := getTime(event1[1])

	beg2 := getTime(event2[0])
	end2 := getTime(event2[1])

	if beg1.Equal(beg2) {
		return true
	}

	if beg1.Before(beg2) {
		return !end1.Before(beg2)
	}

	return !end2.Before(beg1)
}

func getTime(s string) time.Time {
	t, err := time.Parse("15:04", s)
	if err != nil {
		panic(err)
	}

	return t
}
