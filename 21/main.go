package main

import "log"

/*
https://projecteuler.net/problem=21

Let d(n) be defined as the sum of proper divisors
of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are
an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are
1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
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
	const count = 10000

	divSums := make(map[int]int)
	for i := 2; i <= count; i++ {
		divSums[i] = divisorsSum(i)
	}

	var sum int
	for k, v := range divSums {
		if k == v {
			continue
		}

		if divSums[v] == k {
			sum += v + k
			divSums[v] = 0
		}
	}

	log.Printf("The sum of all the amicable numbers under 10000 is %d", sum)
}
