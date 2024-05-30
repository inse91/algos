package _933_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/number-of-recent-calls/description/

func Test(t *testing.T) {
	rc := Constructor()

	assert.Equal(t, 1, rc.Ping(1))
	assert.Equal(t, 2, rc.Ping(100))
	assert.Equal(t, 3, rc.Ping(3001))
	assert.Equal(t, 3, rc.Ping(3002))
}

const gap = 3000

type RecentCounter struct {
	calls []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		calls: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.calls = append(this.calls, t)

	ts := t - gap
	var c int
	for i := len(this.calls) - 1; i > -1; i-- {
		if this.calls[i] < ts {
			break
		}
		c++
	}

	return c
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
