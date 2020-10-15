package main

import (
	"fmt"
)

func main() {
	var n, t int
	fmt.Scanln(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
		t += s[i]
	}
	fmt.Println(t)
}
