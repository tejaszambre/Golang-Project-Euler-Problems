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

func complete(factors []int) bool {
	mul := 1
	for _, v := range factors {
		mul *= v
	}
	if 600851475143 == mul {
		return true
	}
	return false
}

func main() {
	// fmt.Println()
	// fmt.Println(isprime(74))
	Number := 600851475143
	var i = 2
	sqrt := int(math.Sqrt(float64(Number))) + 1
	factors := []int{71, 839, 1471, 6857}
	for {
		if i > sqrt || complete(factors) {
			break
		}
		if Number%i == 0 && isprime(i) {
			factors = append(factors, i)
		}
		fmt.Println(i)
		i++
	}
	fmt.Println(factors[len(factors)-1])
	fmt.Print(complete(factors))

}
