package main

import "log"

/*
https://projecteuler.net/problem=19

You are given the following information,
but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century
unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century
(1 Jan 1901 to 31 Dec 2000)?
*/

func daysInMonth(m, y int) (res int) {
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		res = 31
	case 4, 6, 9, 11:
		res = 30
	case 2:
		res = 28
		if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
			res = 29
		}
	default:
		log.Fatal("invalid logic")
	}

	return
}

const (
	monday = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

type date struct {
	d, m, y int
}

func (d *date) inc() {
	if d.d == daysInMonth(d.m, d.y) {
		d.d = 1

		if d.m == 12 {
			d.m = 1
			d.y++
		} else {
			d.m++
		}
	} else {
		d.d++
	}
}

func main() {
	// find day of week for start
	// given puzzle input
	curr := date{1, 1, 1900}
	dow := monday

	start := date{1, 1, 1901}
	for curr != start {
		dow++
		curr.inc()
	}

	// capture all days of the week until end
	days := make(map[date]int)
	end := date{31, 12, 2000}
	for curr != end {
		days[curr] = dow % 7
		dow++
		curr.inc()
	}

	// count 1st sundays
	var sum int
	for k, v := range days {
		if k.d == 1 && v == sunday {
			sum++
		}
	}

	log.Printf("There are %d Sundays on the first of the month during the twentieth century.", sum)
}
