package _22_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/bits"
	"strconv"
	"strings"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  []string
	}{
		{
			name: "1",
			n:    1,
			res:  []string{"()"},
		},
		{
			name: "3",
			n:    3,
			res:  []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "5",
			n:    5,
			res:  []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, generateParenthesis(c.n))
		})
	}
}

func generateParenthesis(n int) []string {
	const op rune = '('
	const cl rune = ')'

	q := pow(2, n*2)
	res := make([]string, 0, q/8)

f:
	for i := q / 2; i < q; i++ {
		if bits.OnesCount64(uint64(i)) != n {
			continue
		}
		s := strconv.FormatInt(int64(i), 2)

		var cc int
		var sb strings.Builder
		for _, r := range s {
			if r == '0' {
				cc--
				sb.WriteRune(cl)
			} else {
				cc++
				sb.WriteRune(op)
			}
			if cc < 0 {
				continue f
			}

		}

		fmt.Println(sb.String())
		res = append(res, sb.String())
	}

	return res
}

func pow(num, n int) int {
	return int(math.Pow(float64(num), float64(n)))
}
