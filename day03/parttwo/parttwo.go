package parttwo

import (
	"github.com/k-nox/aoc2023/day03/partone"
	"github.com/k-nox/aoc2023/util"
	"strconv"
	"unicode"
)

func Solve() int {
	fscanner := util.NewFileScanner("day03/input/input.txt")
	defer fscanner.Close()
	sum := 0

	grid := partone.ParseGrid(fscanner)

	// loop through grid until a gear is found
	for rowIdx, row := range grid {
		for colIdx, char := range row {
			if !isGear(char) {
				continue
			}

			partNums := findPartNumbers(grid, rowIdx, colIdx)
			if len(partNums) != 2 {
				continue
			}
			ratio := partNums[0] * partNums[1]
			sum += ratio
		}
	}
	return sum
}

func isGear(char rune) bool {
	return string(char) == "*"
}

func findPartNumbers(grid [][]rune, gearRow int, gearCol int) []int {
	var partNumbers []int
	// check above
	if gearRow-1 >= 0 {
		numFound := false
		for col := gearCol - 1; col <= gearCol+1; col++ {
			if col >= 0 && col < len(grid[gearRow-1]) {
				if unicode.IsDigit(grid[gearRow-1][col]) {
					if !numFound {
						// parse the full part number
						partNum := grabPartNumber(grid[gearRow-1], col)
						partNumbers = append(partNumbers, partNum)
						numFound = true
					}
				} else {
					numFound = false
				}
			}
		}
	}
	// check same row
	if unicode.IsDigit(grid[gearRow][gearCol-1]) {
		partNumbers = append(partNumbers, grabPartNumber(grid[gearRow], gearCol-1))
	}

	if unicode.IsDigit(grid[gearRow][gearCol+1]) {
		partNumbers = append(partNumbers, grabPartNumber(grid[gearRow], gearCol+1))
	}
	// check below
	if gearRow+1 < len(grid) {
		numFound := false
		for col := gearCol - 1; col <= gearCol+1; col++ {
			if col >= 0 && col < len(grid[gearRow+1]) {
				if unicode.IsDigit(grid[gearRow+1][col]) {
					if !numFound {
						// parse the full part number
						partNum := grabPartNumber(grid[gearRow+1], col)
						partNumbers = append(partNumbers, partNum)
						numFound = true
					}
				} else {
					numFound = false
				}
			}
		}
	}

	return partNumbers
}

func grabPartNumber(row []rune, foundIdx int) int {
	start := foundIdx
	for i := foundIdx - 1; i >= 0 && unicode.IsDigit(row[i]); i-- {
		start = i
	}
	numStr := ""
	for i := start; i < len(row) && unicode.IsDigit(row[i]); i++ {
		numStr += string(row[i])
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}

	return num
}
