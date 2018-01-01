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

func written(i int) (res string) {
	switch {
	case i == 0:
	case i == 1:
		res = "one"
	case i == 2:
		res = "two"
	case i == 3:
		res = "three"
	case i == 4:
		res = "four"
	case i == 5:
		res = "five"
	case i == 6:
		res = "six"
	case i == 7:
		res = "seven"
	case i == 8:
		res = "eight"
	case i == 9:
		res = "nine"
	case i == 10:
		res = "ten"
	case i == 11:
		res = "eleven"
	case i == 12:
		res = "twelve"
	case i == 13:
		res = "thirteen"
	case i == 14:
		res = "fourteen"
	case i == 15:
		res = "fifteen"
	case i == 16:
		res = "sixteen"
	case i == 17:
		res = "seventeen"
	case i == 18:
		res = "eighteen"
	case i == 19:
		res = "nineteen"
	case i == 20:
		res = "twenty"
	case i == 30:
		res = "thirty"
	case i == 40:
		res = "forty"
	case i == 50:
		res = "fifty"
	case i == 60:
		res = "sixty"
	case i == 70:
		res = "seventy"
	case i == 80:
		res = "eighty"
	case i == 90:
		res = "ninety"
	case i == 1000:
		res = "one thousand"
	case i/1000000 >= 1:
		res = written(i/1000000) + " million " + written(i%1000000)
	case i/1000 >= 1:
		res = written(i/1000) + " thousand " + written(i%1000)
	case i/100 >= 1 && i%100 == 0:
		res = written(i/100) + " hundred"
	case i/100 >= 1 && i%100 != 0:
		res = written(i/100) + " hundred and " + written(i%100)
	case i/10 >= 1 && i%10 > 0:
		res = written(i-i%10) + "-" + written(i%10)
	}

	return
}

func main() {
	const count = 1000

	words := make([]string, 0)
	for i := 1; i <= count; i++ {
		words = append(words, written(i))
	}
	//for _, v := range words {
	//	log.Print(v)
	//}

	var sum int
	for _, v := range words {
		for _, w := range v {
			switch w {
			case ' ':
			case '-':
			default:
				sum++
			}
		}
	}

	log.Printf("There are %d letters used for numbers 1 through %d.", sum, count)
}
