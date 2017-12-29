package main

import "log"

/*
https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by
each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly
divisible by all of the numbers from 1 to 20?
*/

func main() {
	i := 20
out:
	for ; ; i++ {
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				continue out
			}
		}

		break
	}

	log.Printf("The smallest positive number that is evenly divisible by all of the numbers from 1 to 20 is %d", i)
}
