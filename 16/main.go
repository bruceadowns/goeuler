package main

import (
	"log"
	"strconv"
)

/*
https://projecteuler.net/problem=16

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

func main() {
	const exp = 1000

	num := float64(1)
	for i := 0; i < exp; i++ {
		num *= 2
	}
	//log.Printf("num: %.0f", num)

	var sum int32
	s := strconv.FormatFloat(num, 'f', 0, 64)
	for _, v := range s {
		sum += v - '0'
	}

	log.Printf("The sum of the digits of the number 2^1000 is %d.", sum)
}
