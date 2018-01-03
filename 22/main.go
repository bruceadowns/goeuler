package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

/*
https://projecteuler.net/problem=22

Using names.txt (a 46K text file containing over five-thousand first names)
begin by sorting it into alphabetical order.
Then working out the alphabetical value for each name,
multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order,
COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}

	names := make([]string, 0)
	for _, v := range strings.Split(scanner.Text(), ",") {
		if len(v) < 3 || (v[0] != '"' && v[len(v)-1] != '"') {
			log.Fatal("invalid input")
		}

		names = append(names, v[1:len(v)-1])
	}

	sort.Strings(names)

	var sum int
	for k, v := range names {
		var worth int32
		for _, vv := range v {
			worth += vv - 'A' + 1
		}

		sum += (k + 1) * int(worth)
	}

	log.Printf("The total of all the name scores in the file is %d", sum)
}
