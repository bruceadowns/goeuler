package main

import "log"

/*
https://projecteuler.net/problem=14

The following iterative sequence is defined for the set of positive integers:
n -> n/2 (n is even)
n -> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 40 20 10 5 16 8 4 2 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
Although it has not been proved yet (Collatz Problem),
it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

func main() {
	const ceil = 1000000

	var num struct {
		i   int
		max int
	}
	for i := 1; i < ceil; i++ {
		n := i
		var count int
		for n != 1 {
			count++

			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
		}
		count++

		if num.max < count {
			num.max = count
			num.i = i
		}
	}

	log.Printf("The starting number (under %d) produces the longest chain is %d", ceil, num.i)
}
