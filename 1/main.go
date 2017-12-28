package main

import (
	"log"
	"os"
	"strconv"
)

/*
https://projecteuler.net/problem=1

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func main() {
	count := 10
	if len(os.Args) == 2 {
		if n, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(err)
		} else {
			count = n
		}
	}

	sum := 0
	for i := 0; i < count; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	log.Printf("The sum of all the multiples of 3 or 5 below %d is %d", count, sum)
}
