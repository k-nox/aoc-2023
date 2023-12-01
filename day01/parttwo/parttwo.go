package parttwo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Solve() int {
	sum := 0
	// read file one line at a time
	file, err := os.Open("day01/input/input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr := scanner.Text()
		numStr := ""
		// loop through, checking each char to see if it's a digit
		// if it's not, continue looping and checking substrings
		// once the first match is found, break and start again from the back
		i := 0
		numFound := false
		for i < len(curr) && !numFound {
			if unicode.IsDigit(rune(curr[i])) {
				numStr += string(curr[i])
				numFound = true
			}
			for digit, key := range digits {
				if strings.Contains(curr[0:i+1], digit) {
					numStr += key
					numFound = true
					break
				}
			}
			i++
		}

		i = len(curr) - 1
		numFound = false
		for i >= 0 && !numFound {
			if unicode.IsDigit(rune(curr[i])) {
				numStr += string(curr[i])
				numFound = true
				break
			}
			for digit, key := range digits {
				if strings.Contains(curr[i:], digit) {
					numStr += key
					numFound = true
					break
				}
			}
			i--
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		sum += num

	}

	return sum
}
