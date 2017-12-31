package main

import "log"

var cache = make(map[int]uint64)

func fib(n int) uint64 {
	if n < 2 {
		return 1
	}

	if c, ok := cache[n]; ok {
		return c
	}

	res := fib(n-1) + fib(n-2)
	cache[n] = res
	return res
}

func main() {
	for i := 0; i < 50; i++ {
		log.Printf("fib(%d): %d", i, fib(i))
	}
}
