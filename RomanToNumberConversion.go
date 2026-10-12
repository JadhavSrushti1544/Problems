package main

// import "fmt"

// func main() {
// 	var s string
// 	fmt.Println("String: ")
// 	fmt.Scan(&s)

// 	total := 0

// 	for i := 0; i < len(s); i++ {
// 		value := 0

// 		switch s[i] {
// 		case 'I' :
// 			value = 1
// 		case 'V' :
// 			value = 5
// 		case 'X':
// 			value = 10
// 		case 'L':
// 			value = 50
// 		case 'C':
// 			value = 100
// 		case 'D':
// 			value = 500
// 		case 'M':
// 			value = 1000
// 		}

// 		if i+1 < len(s) {
// 			next := 0

// 			switch s[i+1] {
// 			case 'I' :
// 				next = 1
// 			case 'V':
// 				next = 5
// 			case 'X':
// 				next = 10
// 			case 'L':
// 				next = 50
// 			case 'C':
// 				next = 100
// 			case 'D':
// 				next = 500
// 			case 'M':
// 				next = 1000
// 			}

// 			if value < next {
// 				total -= value
// 			} else {
// 				total += value
// 			}
// 		} else {
// 			total += value
// 		}
// 	}

// 	fmt.Println(total)
// }