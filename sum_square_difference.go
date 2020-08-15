package main

import (
	"fmt"
)

func sumofsquare(n int) (sum uint64) {
	sum = uint64((n * (n + 1) * (2*n + 1)) / 6)
	return sum
}

func squareofsum(n int) (square uint64) {
	square = uint64((n * (n + 1)) / 2)
	return square * square
}

func main() {
	fmt.Print("Enter the number: ")
	str := 0
	fmt.Scanln(&str)
	diff := squareofsum(str) - sumofsquare(str)
	fmt.Println(diff)
}
