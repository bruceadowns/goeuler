package main

import (
	"log"
	"strconv"
)

/*
https://projecteuler.net/problem=4

A palindromic number reads the same both ways.
The largest palindrome made from the product
of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func isPalindrome(p int) bool {
	s := strconv.Itoa(p)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	var max int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			p := i * j
			if isPalindrome(p) {
				if max < p {
					max = p
				}
			}
		}
	}

	log.Printf("The largest palindrome made from the product of two 3-digit numbers %d", max)
}
