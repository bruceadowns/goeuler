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

var cache = make(map[coord]int64)

func routes(a coord) int64 {
	if a.x == 0 || a.y == 0 {
		return 1
	}

	if n, ok := cache[a]; ok {
		return n
	}

	r := routes(coord{a.x - 1, a.y}) + routes(coord{a.x, a.y - 1})
	cache[a] = r
	return r
}

func main() {
	count := routes(coord{20, 20})
	log.Printf("There are %d routes through a 20x20 grid.", count)
}
