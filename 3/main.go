package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

/*
https://projecteuler.net/problem=3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func isPrime(i int64) bool {
	for j := int64(2); j < int64(math.Sqrt(float64(i))); j++ {
		if i%j == 0 {
			return false
		}
	}

	return true
}

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	prime := int64(13195)
	if len(os.Args) == 2 {
		if n, err := strconv.ParseInt(os.Args[1], 10, 64); err != nil {
			log.Fatal(err)
		} else {
			prime = n
		}
	}

	primes := make([]int64, 0)
	for i := int64(2); i <= int64(math.Sqrt(float64(prime))); i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	sort.Sort(sort.Reverse(Int64Slice(primes)))

	//max := func() int64 {
	//	for _, v := range primes {
	//		if prime%v == 0 {
	//			return v
	//		}
	//	}
	//	return -1
	//}()

	var max int64
	for _, v := range primes {
		if prime%v == 0 {
			max = v
			break
		}
	}

	log.Printf("The largest prime factor of %d is %d", prime, max)
}
