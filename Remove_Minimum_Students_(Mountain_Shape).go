
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	parts := strings.Split(input, ",")
	n := len(parts)

	h := make([]int, n)

	for i := 0; i < n; i++ {
		h[i], _ = strconv.Atoi(parts[i])
	}

	inc := make([]int, n)
	dec := make([]int, n)

	for i := 0; i < n; i++ {
		inc[i] = 1

		for j := 0; j < i; j++ {
			if h[j] < h[i] && inc[j]+1 > inc[i] {
				inc[i] = inc[j] + 1
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		dec[i] = 1

		for j := i + 1; j < n; j++ {
			if h[j] < h[i] && dec[j]+1 > dec[i] {
				dec[i] = dec[j] + 1
			}
		}
	}

	longest := 0

	for i := 1; i < n-1; i++ {
		if inc[i] > 1 && dec[i] > 1 {
			length := inc[i] + dec[i] - 1

			if length > longest {
				longest = length
			}
		}
	}

	if longest == 0 {
		fmt.Println("NOT POSSIBLE")
	} else {
		fmt.Println(n - longest)
	}
}
