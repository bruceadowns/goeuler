package main

import (
	"log"
	"sort"
	"time"
)

/*
https://projecteuler.net/problem=23

A perfect number is a number for which the sum of
its proper divisors is exactly equal to the number.

For example, the sum of the proper divisors of 28
would be 1 + 2 + 4 + 7 + 14 = 28,
which means that 28 is a perfect number.

A number n is called deficient if the sum of its
proper divisors is less than n and it is called
abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16,
the smallest number that can be written
as the sum of two abundant numbers is 24.

By mathematical analysis, it can be shown that all
integers greater than 28123 can be written as the
sum of two abundant numbers.
However, this upper limit cannot be reduced any
further by analysis even though it is known that
the greatest number that cannot be expressed as
the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

func divisorsSum(i int) (res int) {
	res = 1
	for j := 2; j <= i/2; j++ {
		if i%j == 0 {
			res += j
		}
	}

	return
}

func main() {
	const largest = 28123

	start := time.Now().UnixNano()

	// make list of divisor sums
	divSums := make(map[int]int)
	for i := 1; i <= largest; i++ {
		divSums[i] = divisorsSum(i)
	}

	// make list of abundant numbers
	abundants := make([]int, 0)
	for k, v := range divSums {
		if k < v {
			abundants = append(abundants, k)
		}
	}

	// make list of abundant sum combinations
	sort.Ints(abundants)
	abundantSums := make(map[int]struct{}, 0)
	for i := 0; i < len(abundants); i++ {
		for j := i; j < len(abundants); j++ {
			sum := abundants[i] + abundants[j]
			if sum > 28123 {
				break
			}
			abundantSums[sum] = struct{}{}
		}
	}

	// count positive numbers not in combo list
	var sum, count int
	for i := 1; i <= largest; i++ {
		if _, ok := abundantSums[i]; !ok {
			sum += i
			count++
		}
	}

	end := (time.Now().UnixNano() - start) / 1000000

	log.Printf("The sum of all the positive integers (%d) which cannot be written as the sum of two abundant numbers is %d (duration: %d)", count, sum, end)
}
