package main

import (
	"log"
	"math/big"
)

/*
https://projecteuler.net/problem=20

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

func fac(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	res := big.NewInt(n)
	return res.Mul(res, fac(n-1))
}

func main() {
	const factor = 100
	f := fac(factor)

	var sum int32
	for _, v := range f.String() {
		sum += v - '0'
	}

	log.Printf("The sum of the digits in the number %d! is %d.", factor, sum)
}
