package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringtonumberslice(s string) (a []int) {
	b := []string{}
	b = strings.Split(s, "")
	for _, v := range b {
		if v, err := strconv.Atoi(v); err == nil {
			a = append(a, v)
		}
	}
	return
}

func ispalindrome(no int) bool {
	mul := stringtonumberslice(strconv.Itoa(no))
	lenMul := len(mul)
	for i := 0; i < lenMul/2; i++ {
		if mul[i] != mul[lenMul-i-1] {
			return false
		}
	}
	return true
}

func islargestpalindrome(no1 int, no2 int) (mul int) {
	max := 0
	for i := no1; i > 100; i-- {
		for j := no2; j > 100; j-- {
			mul = i * j
			if ispalindrome(mul) {
				if mul > max {
					max = mul
					fmt.Println(i)
					fmt.Println(j)
				}
			}
		}
	}
	return max
}

func main() {
	no1 := 999
	no2 := 999
	fmt.Println(islargestpalindrome(no1, no2))
}
