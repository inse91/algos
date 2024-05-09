package _884_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		s1   string
		s2   string
		res  []string
	}{
		{
			name: "1",
			s1:   "this apple is sweet",
			s2:   "this apple is sour",
			res:  []string{"sweet", "sour"},
		},
		{
			name: "2",
			s1:   "apple apple",
			s2:   "banana",
			res:  []string{"banana"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, uncommonFromSentences(c.s1, c.s2))
		})
	}
}

func uncommonFromSentences(s1 string, s2 string) []string {
	res := make([]string, 0)

	m1 := mMap(s1)
	m2 := mMap(s2)

	for k, v := range m1 {
		if !v {
			continue
		}

		_, ok := m2[k]
		if ok {
			continue
		}

		res = append(res, k)
	}

	for k, v := range m2 {
		if !v {
			continue
		}

		_, ok := m1[k]
		if ok {
			continue
		}

		res = append(res, k)
	}

	return res
}

func mMap(s string) map[string]bool {
	m := make(map[string]bool)

	words := strings.Split(s, " ")
	for _, w := range words {
		_, ok := m[w]
		if !ok {
			m[w] = true
			continue
		}

		m[w] = false
	}

	return m
}
