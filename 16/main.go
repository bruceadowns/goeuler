package main

import (
	"fmt"
	"log"
	"math"
)

/*
https://projecteuler.net/problem=16

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

func main() {
	var sum int32
	for _, v := range fmt.Sprintf("%.0f", math.Pow(2, 1000)) {
		sum += v - '0'
	}

	log.Printf("The sum of the digits of the number 2^1000 is %d.", sum)
}
