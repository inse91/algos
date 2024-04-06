package _682_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestBaseball(t *testing.T) {
	cases := []struct {
		name string
		ops  []string
		res  int
	}{
		{
			name: "1",
			ops:  []string{"5", "2", "C", "D", "+"},
			res:  30,
		},
		{
			name: "2",
			ops:  []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			res:  27,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := calPoints(c.ops)
			assert.Equal(t, c.res, res)
		})
	}
}

func calPoints(operations []string) int {
	var res int
	recs := make([]int, 0, len(operations))

	for _, op := range operations {
		switch op {
		case "+":
			last := recs[len(recs)-1]
			plast := recs[len(recs)-2]
			recs = append(recs, last+plast)
		case "D":
			last := recs[len(recs)-1]
			recs = append(recs, last*2)
		case "C":
			recs = recs[:len(recs)-1]
		default:
			r, err := strconv.Atoi(op)
			if err != nil {
				panic(err)
			}
			recs = append(recs, r)
		}
	}

	for _, r := range recs {
		res += r
	}

	return res
}

//for _, op := range operations {
//switch op {
//case "+":
//newRec := prev + pprev
//res += newRec
//pprev, prev = prev, newRec
////prev, pprev = newRec, prev
//case "D":
//newRec := prev * 2
//res += newRec
//pprev, prev = prev, newRec
////prev, pprev = newRec, prev
//case "C":
//newRec := -prev
//res += newRec
////prev, pprev = newRec, prev
//pprev, prev = prev, newRec
//default:
//newRec, err := strconv.Atoi(op)
//if err != nil {
//panic(err)
//}
//res += newRec
//pprev, prev = prev, newRec
////prev, pprev = newRec, prev
//}
//}
