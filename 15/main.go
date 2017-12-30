package main

import "log"

/*
https://projecteuler.net/problem=15

Starting in the top left corner of a 2×2 grid,
and only being able to move to the right and down,
there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
*/

type coord struct {
	x, y int
}

var cache = make(map[coord]int)

func routes(x, y int) int {
	if x == 0 || y == 0 {
		return 1
	}

	c := coord{x, y}
	if n, ok := cache[c]; ok {
		return n
	}

	r := routes(x-1, y) + routes(x, y-1)
	cache[c] = r
	return r
}

func main() {
	count := routes(20, 20)
	log.Printf("There are %d routes through a 20x20 grid.", count)
}
