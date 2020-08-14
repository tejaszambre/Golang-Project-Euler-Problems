package main

import (
	"fmt"
)

func multiple20(no int, r int) bool {
	for i := 1; i <= r; i++ {
		if no%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Hello")
	r := 0
	fmt.Scanln(&r)
	no := r
	for {
		if multiple20(no, r) {
			break
		}
		no += r
	}
	fmt.Println(no)
}
