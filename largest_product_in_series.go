package main

/* The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.

7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450

Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?
*/

import (
	"fmt"
	"strconv"
	"strings"
)

var numbers = []int{}
var max uint64 = 0

func maxmultiple(startindex int, endindex int) {
	var mul uint64 = 1
	for i := startindex; i < endindex; i++ { // multiple the 13 digits
		mul *= uint64(numbers[i])
	}
	if mul > max { // if max then update
		max = mul
	}
}

func main() {
	str := ""
	fmt.Scan(&str)
	s := strings.Split(str, "")

	for _, v := range s { // convert the slice of string to slice of int
		if no, err := strconv.Atoi(v); err == nil {
			numbers = append(numbers, no)
		}
	}

	startindex := 0
	endindex := 0
	for i, v := range numbers {
		if v == 0 { // check if the number is zero
			startindex = i + 1 // if yes update the start and end index to index+1
			endindex = i + 1
		} else {
			endindex += 1                  // else update the end index to index+1
			if endindex-startindex == 13 { // check we have got 13 consecutive digit
				maxmultiple(startindex, endindex)
				startindex += 1 // update the start index by adding one in it, to maintain 13 lenght
			}
		}
	}
	fmt.Println(max)

}
