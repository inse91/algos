package _1598_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/crawler-log-folder/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		logs []string
		res  int
	}{
		{
			name: "1",
			logs: []string{"d1/", "d2/", "../", "d21/", "./"},
			res:  2,
		},
		{
			name: "2",
			logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			res:  3,
		},
		{
			name: "3",
			logs: []string{"d1/", "../", "../", "../"},
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minOperations(c.logs))
		})
	}
}

func minOperations(logs []string) int {
	const (
		up    = "../"
		stays = "./"
	)

	var steps int
	for _, log := range logs {
		if log == stays {
			continue
		}

		if log != up {
			steps++
			continue
		}

		if steps != 0 {
			steps--
		}
	}

	return steps
}
