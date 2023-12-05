package partone

import (
	"github.com/k-nox/aoc2023/util"
	"strconv"
	"unicode"
)

func Solve() int {
	fscanner := util.NewFileScanner("day03/input/input.txt")
	defer fscanner.Close()
	sum := 0
	grid := ParseGrid(fscanner)

	// loop through the grid and keep track of the current num string we are looking at
	for rowIdx, row := range grid {
		numStr := ""
		isPartNum := false
		for colIdx, char := range row {
			if unicode.IsDigit(char) {
				numStr += string(char)
				// if the current char is a num, check all the adjacent sides to see if there is a symbol
				if !isPartNum {
					// if it is, then mark the num as a part number and parse it once we find the next char that is not a num
					isPartNum = IsPartNumber(grid, rowIdx, colIdx)
				}

				// process numStr only if we are at the end of the line or the next digit is not a num
				if colIdx+1 == len(row) || !unicode.IsDigit(row[colIdx+1]) {
					if numStr != "" {
						if isPartNum {
							num, err := strconv.Atoi(numStr)
							if err != nil {
								panic(err)
							}
							sum += num
							isPartNum = false
						}
						// then reset the numstr and keep going
						numStr = ""
					}
				}
			}
		}
	}

	return sum
}

func ParseGrid(scanner *util.FileScanner) [][]rune {
	var grid [][]rune
	for scanner.Scan() {
		// read the current line
		line := scanner.Text()
		// add each rune to the grid row
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		// add the row to the grid
		grid = append(grid, row)
	}
	return grid
}

func IsPartNumber(grid [][]rune, rowIdx int, colIdx int) bool {
	// check above
	if rowIdx-1 >= 0 {
		for i := colIdx - 1; i <= colIdx+1; i++ {
			if i >= 0 && i < len(grid[rowIdx-1]) {
				if IsSybmol(grid[rowIdx-1][i]) {
					return true
				}
			}
		}
	}
	// check same row
	for i := colIdx - 1; i <= colIdx+1; i++ {
		if i >= 0 && i < len(grid[rowIdx]) && i != colIdx {
			if IsSybmol(grid[rowIdx][i]) {
				return true
			}
		}
	}

	// check below
	if rowIdx+1 < len(grid) {
		for i := colIdx - 1; i <= colIdx+1; i++ {
			if i >= 0 && i < len(grid[rowIdx+1]) {
				if IsSybmol(grid[rowIdx+1][i]) {
					return true
				}
			}
		}
	}
	return false
}

func IsSybmol(char rune) bool {
	if unicode.IsDigit(char) || string(char) == "." {
		return false
	}
	return !unicode.IsDigit(char) && !IsPeriod(char)
}

func IsPeriod(char rune) bool {
	return string(char) == "."
}
