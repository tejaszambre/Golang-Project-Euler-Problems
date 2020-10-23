package main

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

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

func problem10() (answer uint64) {
	str := 0
	answer = 5
	fmt.Println("Enter the number :")
	fmt.Scanln(&str)
	if str == 1 || str == 2 {
		answer = 0
		return answer
	} else if str == 3 {
		answer = 2
		return answer
	}
	for i := 3; i < str; i++ {
		if i%2 != 0 && i%3 != 0 && isprime(i) {
			answer += uint64(i)
		}
	}
	return answer
}

func main() {
	answer := problem10()
	fmt.Println(answer)
}
