package main

import (
	"bytes"
	"log"
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

// adaptation of SEPA Algorithm
func permute(iterable []int) bool {
	l := len(iterable)
	key := l - 1
	newkey := l - 1

	for (key > 0) && (iterable[key] <= iterable[key-1]) {
		key--
	}

	key--

	if key < 0 {
		return false
	}

	for (newkey > key) && (iterable[newkey] <= iterable[key]) {
		newkey--
	}

	iterable[key], iterable[newkey] = iterable[newkey], iterable[key]

	l--
	key++

	for l > key {
		iterable[l], iterable[key] = iterable[key], iterable[l]
		key++
		l--
	}

	return true
}

func main() {
	iterable := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 999999; i++ {
		permute(iterable)
	}

	sb := bytes.Buffer{}
	for _, v := range iterable {
		sb.WriteRune('0' + rune(v))
	}

	log.Printf("The millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9 is %s.", sb.String())
}
