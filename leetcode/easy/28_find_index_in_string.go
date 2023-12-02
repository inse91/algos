package easy

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func Run28() {
	fmt.Println(strStr("sadbutsad", "sad"))
}
