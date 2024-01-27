package _459__test

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "abab",
			res:  true,
		},
		{
			name: "2",
			s:    "aba",
			res:  false,
		},
		{
			name: "3",
			s:    "abcabcabcabc",
			res:  true,
		},
		{
			name: "4",
			s:    "a",
			res:  false,
		},
		{
			name: "5",
			s:    "aaa",
			res:  true,
		},
		{
			name: "6",
			s:    "ababba",
			res:  false,
		},
		{
			name: "7",
			s:    "ababab",
			res:  true,
		},
		{
			name: "8",
			s:    "abaaba",
			res:  true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := repeatedSubstringPattern(c.s)
			if res != c.res {
				t.Fatalf("expected %v, got %v", c.res, res)
			}
		})
	}

}

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	lenS := len(s)

	for divisor := 2; divisor <= lenS; divisor++ {
		if lenS%divisor != 0 {
			continue
		}

		parts := divide(s, divisor)
		if len(parts) == 0 {
			continue
		}
		sample := parts[0]
		allPartsAreTheSame := true
		for _, part := range parts {
			if part != sample {
				allPartsAreTheSame = false
				break
			}
		}
		if allPartsAreTheSame {
			return true
		}
	}

	return false
}

func divide(s string, numOfPcs int) []string {
	parts := make([]string, 0, numOfPcs)
	partLen := len(s) / numOfPcs
	for i := 0; i < numOfPcs; i++ {
		start := i * partLen
		end := start + partLen
		parts = append(parts, s[start:end])
	}
	return parts
}
