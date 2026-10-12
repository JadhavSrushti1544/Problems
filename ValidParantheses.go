
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	stack := []byte{}
	valid := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				valid = false
				break
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if ch == ')' && last != '(' {
				valid = false
			} else if ch == '}' && last != '{' {
				valid = false
			} else if ch == ']' && last != '[' {
				valid = false
			}

			if !valid {
				break
			}
		}
	}

	if len(stack) != 0 {
		valid = false
	}

	fmt.Println(valid)
}
