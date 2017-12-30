package main

import (
	"log"
	"time"
)

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

var size = coord{20, 20}

var cache = make(map[coord]int64)

func routes(a coord) (res int64) {
	if a == size {
		return 1
	}

	if a.x < size.x {
		next := coord{a.x + 1, a.y}
		if n, ok := cache[next]; ok {
			res += n
		} else {
			subRoutes := routes(next)
			cache[next] = subRoutes
			res += subRoutes
		}
	}

	if a.y < size.y {
		next := coord{a.x, a.y + 1}
		if n, ok := cache[next]; ok {
			res += n
		} else {
			subRoutes := routes(next)
			cache[next] = subRoutes
			res += subRoutes
		}
	}

	return
}

func main() {
	startingTime := time.Now().UnixNano()
	count := routes(coord{0, 0})
	dur := (time.Now().UnixNano() - startingTime) / 1000000

	log.Printf(
		"There are %d routes through a %dx%d grid. [duration: %dms]",
		count, size.x, size.y, dur)
}
