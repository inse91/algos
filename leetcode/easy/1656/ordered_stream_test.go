package _1656_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/design-an-ordered-stream/description/

func Test(t *testing.T) {
	os := Constructor(5)

	require.Equal(t, []string{}, os.Insert(3, "ccccc"))
	require.Equal(t, []string{"aaaaa"}, os.Insert(1, "aaaaa"))
	require.Equal(t, []string{"bbbbb", "ccccc"}, os.Insert(2, "bbbbb"))
	require.Equal(t, []string{}, os.Insert(5, "eeeee"))
	require.Equal(t, []string{"ddddd", "eeeee"}, os.Insert(4, "ddddd"))
}

type OrderedStream struct {
	store []string
	q     int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		store: make([]string, n),
	}
}

func (os *OrderedStream) Insert(idKey int, value string) []string {
	st := os.q
	os.store[idKey-1] = value

	for {
		if os.store[os.q] == "" {
			break
		}

		if os.q == len(os.store)-1 {
			return os.store[st:]
		}

		os.q++
	}

	return os.store[st:os.q]
}
