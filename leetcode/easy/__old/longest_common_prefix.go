package __old

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) (res string) {
	l := len(strs)
	if l == 0 {
		return
	}
	orig := strs[0]
	if l == 1 {
		return orig
	}

	strs = strs[1:]
	for i := range orig {
		for _, s := range strs {
			if i >= len(s) || orig[i] != s[i] {
				return
			}
		}
		res += string(orig[i])
	}
	return
}

func longestCommonPrefixOld(strs []string) (res string) {
	minLen := 200
	for _, s := range strs {
		minLen = int(math.Min(float64(minLen), float64(len(s))))
	}
	fmt.Println(minLen)

	res = ""
	for i := 0; i < int(minLen); i++ {
		A := (strs[0])[i]
		for _, s := range strs {
			B := s[i]
			if B != A {
				return
			}
		}
		res += string(A)
	}
	return
}

func RunLongestCommonPrefix() {
	s := []string{"flower", "flo", "flower"}
	fmt.Println(longestCommonPrefix(s))
}
