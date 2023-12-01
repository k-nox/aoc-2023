package partone

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

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
		// loop through string front-to-back looking for a digit
		for _, char := range curr {
			if unicode.IsDigit(char) {
				numStr += string(char)
				break
			}
		}

		for i := len(curr) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(curr[i])) {
				numStr += string(curr[i])
				break
			}
		}

		// convert numStr to num and add it to sum
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		sum += num
	}

	// return sum
	return sum
}
