package main

import (
	"log"
	"math"
)

/*
https://projecteuler.net/problem=7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

func isPrime(i int) bool {
	for j := 2; j <= int(math.Floor(math.Sqrt(float64(i)))); j++ {
		if i%j == 0 {
			return false
		}
	}

	return true
}

func main() {
	primeCount := 0
	count := 2
	for {
		if isPrime(count) {
			primeCount++

			if primeCount == 10001 {
				break
			}
		}

		count++
	}

	log.Printf("The %d prime number is %d", primeCount, count)
}
