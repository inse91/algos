package __old

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	ln := len(arr)
	if ln == 0 {
		return 0
	}

	//fmt.Printf("arr: %#v\n", arr)
	//lastNonEmptyString := arr[ln-1]
	//for c := 1; lastNonEmptyString == "" && ln-1-c >= 0; c++ {
	//	lastNonEmptyString = arr[ln-1-c]
	//}
	return len(arr[ln-1])
}

func Run58() {
	//input := "Hello world"
	input := "   fly me   to   the moon  "
	res := lengthOfLastWord(input)
	fmt.Printf("res: %d\n", res)
}
