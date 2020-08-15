package main

import (
	"fmt"
	"math"
)

func isprime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	str := 0
	primenumber := 2
	fmt.Println("Enter the number :")
	fmt.Scanln(&str)
	count := 2
	if str == 1 || str == 2 {
		goto Print
	}
	for i := 4; i < str*20; i++ {
		if i%2 != 0 && i%3 != 0 && isprime(i) {
			count += 1
			if count == str {
				primenumber = i
				break
			}
		}
	}
Print:
	if str == 2 {
		primenumber = 3
	}
	fmt.Println(primenumber)
}
