package _2037_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-out-of-three/description/

func Test(t *testing.T) {
	cases := []struct {
		name     string
		seats    []int
		students []int
		res      int
	}{
		{
			name:     "1",
			seats:    []int{3, 1, 5},
			students: []int{2, 7, 4},
			res:      4,
		},
		{
			name:     "2",
			seats:    []int{4, 1, 5, 9},
			students: []int{1, 3, 2, 6},
			res:      7,
		},
		{
			name:     "3",
			seats:    []int{2, 2, 6, 6},
			students: []int{1, 3, 2, 6},
			res:      4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minMovesToSeat(c.seats, c.students))
		})
	}
}

func minMovesToSeat(seats []int, students []int) int {
	slices.Sort(seats)
	slices.Sort(students)

	var res int
	for i := range seats {
		res += abs(seats[i] - students[i])
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
