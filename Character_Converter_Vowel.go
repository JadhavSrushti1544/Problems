 package main

// import "fmt"

// func main() {
// 	var s string
// 	fmt.Scan(&s)

// 	result := ""

// 	for i := 0; i < len(s); i++ {
// 		ch := s[i]

// 		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
// 			count := i + 1

// 			if count > 5 {
// 				count = 5
// 			}

// 			for j := 0; j < count; j++ {
// 				result += "_"
// 			}
// 		} else {
// 			result += string(ch)
// 		}
// 	}

// 	fmt.Println(result)
// }