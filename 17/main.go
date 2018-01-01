package main

import "log"

/*
https://projecteuler.net/problem=17

If the numbers 1 to 5 are written out in words:
one, two, three, four, five,
thenthere are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive
were written out in words,
how many letters would be used?
*/

func main() {
	ones := len("onetwothreefourfivesixseveneightnine")
	log.Printf("ones: %d", ones)

	teens := len("teneleventwelvethirteenfourteenfifteensixteenseventeeneighteennineteen")
	log.Printf("teens: %d", teens)

	tens := len("twenty")*10 + ones
	tens += len("thirty")*10 + ones
	tens += len("forty")*10 + ones
	tens += len("fifty")*10 + ones
	tens += len("sixty")*10 + ones
	tens += len("seventy")*10 + ones
	tens += len("eighty")*10 + ones
	tens += len("ninety")*10 + ones
	log.Printf("tens: %d", tens)

	totalones := ones * 100
	log.Printf("totalones: %d", totalones)

	under100 := ones + teens + tens
	log.Printf("under100: %d", under100)

	totalunder100 := under100 * 9
	log.Printf("totalunder100: %d", totalunder100)

	hundred := 9 * len("hundred")
	log.Printf("hundred: %d", hundred)

	hundredand := len("hundredand") * 891
	log.Printf("hundredand: %d", hundredand)

	thousands := len("onethousand")
	log.Printf("thousands: %d", thousands)

	sum := under100 + totalones + totalunder100 + hundred + hundredand + thousands

	log.Printf("There are %d letters used.", sum)
}
