package main

import (
	"log"
	"math/big"
)

/*
https://projecteuler.net/problem=25

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

var cache = make(map[int]*big.Int)

func fib(n int) *big.Int {
	if n < 2 {
		return big.NewInt(1)
	}

	if c, ok := cache[n]; ok {
		return c
	}

	// copy first fib; big.Add mangles the cache value
	f := &big.Int{}
	f.Set(fib(n - 1))
	f.Add(f, fib(n-2))

	cache[n] = f

	return f
}

func main() {
	const digitLen = 1000

	var idx int
	for {
		f := fib(idx)

		if len(f.String()) >= digitLen {
			break
		}

		idx++
	}

	log.Printf("The index of the first term in the Fibonacci sequence to contain %d digits is %d", digitLen, idx+1)
}
