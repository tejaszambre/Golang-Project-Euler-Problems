package main

import "fmt"

var memoize = map[int]uint64{}

func fibonacci(n int) uint64 {

	_, ok := memoize[n]
	if ok {
		return memoize[n]
	}
	if n == 0 || n == 1 {
		memoize[n] = 1
		return memoize[n]
	}
	memoize[n] = fibonacci(n-1) + fibonacci(n-2)
	return memoize[n]
}

func main() {
	fibonacci(1000)
	var sum uint64 = 0
	for _, v := range memoize {
		if v < 4000000 && v%2 == 0 {
			sum += v
		}
	}
	fmt.Println(sum)
}
