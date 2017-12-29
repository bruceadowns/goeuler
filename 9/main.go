package main

import "log"

/*
https://projecteuler.net/problem=9

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func main() {
	var a, b, c int
out:
	for a = 1; a < 1000; a++ {
		for b = a; b < 1000; b++ {
			for c = b; c < 1000; c++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					break out
				}
			}
		}
	}

	log.Printf("The product of the Pythagorean triplet for which a + b + c = 1000 is %d", a*b*c)
}
