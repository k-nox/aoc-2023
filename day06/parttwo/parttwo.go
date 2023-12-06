package parttwo

import (
	"github.com/k-nox/aoc2023/day06/partone"
	"github.com/k-nox/aoc2023/util"
	"strconv"
	"strings"
)

func Solve() int {
	f := util.NewScannerForInput(6, false)
	defer f.Close()

	race := parseInput(f)

	return race.FindNumberOfPossibleWins()
}

func parseInput(f *util.FileScanner) partone.Race {
	race := partone.Race{}

	for f.Scan() {
		line := f.Text()
		numStr := strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		if race.Time == 0 {
			// we are parsing the time
			race.Time = num
		} else {
			race.Distance = num
		}
	}
	return race
}
