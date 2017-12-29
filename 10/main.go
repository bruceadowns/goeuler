package main

import (
	"log"
	"math"
)

/*
https://projecteuler.net/problem=10

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
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
	primes := make([]int, 0)
	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	var sum int
	for _, v := range primes {
		sum += v
	}

	log.Printf("The sum of all the primes below two million is %d", sum)
}
