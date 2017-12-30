package main

import "log"

func fib(n int) int {
	if n < 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	for i := 0; i < 10; i++ {
		log.Printf("fib(%d): %d", i, fib(i))
	}
}
