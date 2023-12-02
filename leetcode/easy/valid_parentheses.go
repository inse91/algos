package easy

import "fmt"

func isValid(s string) bool {
	//parenthMap := make(map[rune]int)

	//openedMap := make(map[rune]bool)
	//closedMap := make(map[rune]bool)

	opened := make([]rune, 0, len(s))

	for _, ch := range s {

		if ch == '(' || ch == '{' || ch == '[' {
			opened = append(opened, ch)
			continue
		}

		opLen := len(opened)
		if opLen == 0 {
			return false
		}
		lastOpened := opened[opLen-1]

		if ch == ')' && lastOpened != '(' {
			return false
		}
		if ch == ']' && lastOpened != '[' {
			return false
		}
		if ch == '}' && lastOpened != '{' {
			return false
		}
		opened = opened[:opLen-1]
	}

	return len(opened) == 0

}

func RunValidParenth() {
	//s := "([){}[]]"
	s := "("
	fmt.Println(isValid(s))
}
