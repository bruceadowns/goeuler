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

func appendPrime(in []int, i int) []int {
	for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
		if i%j == 0 {
			return in
		}
	}

	return append(in, i)
}

func main() {
	primes := make([]int, 0)
	for i := 2; i < 2000000; i++ {
		primes = appendPrime(primes, i)
	}

	var sum int
	for _, v := range primes {
		sum += v
	}

	log.Printf("The sum of all the primes below two million is %d", sum)
}
