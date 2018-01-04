package main

import (
	"bytes"
	"log"

	"github.com/bruceadowns/goitertools/itertools"
)

/*
https://projecteuler.net/problem=24

A permutation is an ordered arrangement of objects.
For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically,
we call it lexicographic order.

The lexicographic permutations of 0, 1 and 2 are:
012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

func main() {
	p := itertools.Permutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)

	sb := bytes.Buffer{}
	for _, v := range p[999999] {
		sb.WriteRune('0' + rune(v))
	}

	log.Printf("The millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9 is %s.", sb.String())
}
