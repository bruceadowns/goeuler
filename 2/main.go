package main

import (
	"log"
	"os"
	"strconv"
)

/*
https://projecteuler.net/problem=2

Each new term in the Fibonacci sequence is generated by
adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence
whose values do not exceed four million,
find the sum of the even-valued terms.
*/

func main() {
	highWater := 10
	if len(os.Args) == 2 {
		if n, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(err)
		} else {
			highWater = n
		}
	}

	var curr, prev int

	curr = 1
	sum := 0
	for i := 0; curr < highWater; i++ {
		if curr%2 == 0 {
			sum += curr
		}

		prev, curr = curr, curr+prev
	}

	log.Printf("The sum of the even-valued terms below %d is %d", highWater, sum)
}
