package _1880_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/description/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		word1  string
		word2  string
		target string
		res    bool
	}{
		{
			name:   "1",
			word1:  "acb",
			word2:  "cba",
			target: "cdb",
			res:    true,
		},
		{
			name:   "2",
			word1:  "aaa",
			word2:  "a",
			target: "aab",
			res:    false,
		},
		{
			name:   "3",
			word1:  "aaa",
			word2:  "a",
			target: "aaaa",
			res:    true,
		},
		{
			name:   "73",
			word1:  "j",
			word2:  "j",
			target: "bi",
			res:    true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isSumEqual(c.word1, c.word2, c.target))
		})
	}
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var (
		sum1, sum2, sum3 int
		l1, l2, l3       = len(firstWord), len(secondWord), len(targetWord)
	)

	for i, r := range firstWord {
		inc := int(r - 'a')
		for range l1 - 1 - i {
			inc *= 10
		}
		sum1 += inc
	}
	for i, r := range secondWord {
		inc := int(r - 'a')
		for range l2 - 1 - i {
			inc *= 10
		}
		sum2 += inc
	}
	for i, r := range targetWord {
		inc := int(r - 'a')
		for range l3 - 1 - i {
			inc *= 10
		}
		sum3 += inc
	}

	return sum1+sum2 == sum3
}
