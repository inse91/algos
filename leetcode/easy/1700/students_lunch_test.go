package _1700_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/description/

func Test(t *testing.T) {
	cases := []struct {
		name       string
		students   []int
		sandwiches []int
		res        int
	}{
		{
			name:       "1",
			students:   []int{1, 1, 0, 0},
			sandwiches: []int{0, 1, 0, 1},
			res:        0,
		},
		{
			name:       "2",
			students:   []int{1, 1, 1, 0, 0, 1},
			sandwiches: []int{1, 0, 0, 0, 1, 1},
			res:        3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countStudents(c.students, c.sandwiches))
		})
	}
}

func countStudents(students []int, sandwiches []int) int {
	var s0, s1 int
	for _, st := range students {
		if st == 0 {
			s0++
		} else {
			s1++
		}
	}

	for _, food := range sandwiches {
		switch food {
		case 0:
			if s0 == 0 {
				return s1
			}
			s0--
		case 1:
			if s1 == 0 {
				return s0
			}
			s1--
		}
	}

	return 0
}
