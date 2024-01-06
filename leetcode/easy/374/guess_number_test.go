package _274_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const gap = 100

var num int

func TestGuessNumber(t *testing.T) {
	num = 1 + rand.Intn(gap)
	//num = 100
	fmt.Println(num)
	assert.Equal(t, num, guessNumber(gap))
}

func guessNumber(n int) int {
	pick := 1 + rand.Intn(n)
	start := 1
	for {
		guessRes := guess(pick)
		switch guessRes {
		case 0:
			return pick
		case -1:
			start = pick
			pick = start + (n-pick)/2 + 1
		case 1:
			n = pick
			pick = start + (n-start)/2
		}
	}
}

func guess(pick int) int {
	if num > pick {
		return -1
	}
	if num < pick {
		return 1
	}
	return 0
}
