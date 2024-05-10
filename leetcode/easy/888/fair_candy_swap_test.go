package _888_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name       string
		aliceSizes []int
		bobSizes   []int
		res        []int
	}{
		{
			name:       "1",
			aliceSizes: []int{1, 1},
			bobSizes:   []int{2, 2},
			res:        []int{1, 2},
		},
		{
			name:       "2",
			aliceSizes: []int{1, 2},
			bobSizes:   []int{2, 3},
			res:        []int{1, 2},
		},
		{
			name:       "3",
			aliceSizes: []int{2},
			bobSizes:   []int{1, 3},
			res:        []int{2, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res[0], fairCandySwap(c.aliceSizes, c.bobSizes)[0])
			assert.Equal(t, c.res[1], fairCandySwap(c.aliceSizes, c.bobSizes)[1])
		})
	}
}

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	alcSum := sum(aliceSizes)
	bobSum := sum(bobSizes)

	slices.Sort(aliceSizes)
	slices.Sort(bobSizes)

	diff := alcSum - bobSum

	for _, a := range aliceSizes {
		for _, b := range bobSizes {
			d := (a - b) * 2
			if d < diff {
				break
			}
			if d == diff {
				return []int{a, b}
			}
		}
	}

	return nil
}

func sum(in []int) int {
	var s int
	for _, n := range in {
		s += n
	}

	return s
}
