package _1507_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
	"unicode"
)

// https://leetcode.com/problems/reformat-date/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "20th Oct 2052",
			res:  "2052-10-20",
		},
		{
			name: "2",
			s:    "6th Jun 1933",
			res:  "1933-06-06",
		},
		{
			name: "3",
			s:    "26th May 1960",
			res:  "1960-05-26",
		},
		{
			name: "12",
			s:    "22nd Apr 2023",
			res:  "2023-04-22",
		},
		{
			name: "17",
			s:    "1st Apr 2023",
			res:  "2023-04-01",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, reformatDate(c.s))
		})
	}
}

func reformatDate(date string) string {
	const layout = "2 Jan 2006"
	const space = " "

	parts := strings.Split(date, space)
	if len(parts) == 0 {
		panic("no parts")
	}

	parts[0] = trim(parts[0])
	date = strings.Join(parts, space)

	d, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	return d.Format(time.DateOnly)
}

func trim(s string) string {
	for i, r := range s {
		if unicode.IsDigit(r) {
			continue
		}

		return s[:i]
	}

	return s
}
