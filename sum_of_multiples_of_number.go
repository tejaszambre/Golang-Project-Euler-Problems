package main

import (
	"fmt"
)

func main() {
	a := []int{}
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			a = append(a, i)
		}
	}
	for i := range a {
		sum += a[i]
	}

	fmt.Println(sum)
	fmt.Println(a)
}
