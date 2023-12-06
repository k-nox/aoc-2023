package partone

import (
	"github.com/k-nox/aoc2023/util"
	"strconv"
	"strings"
)

func Solve() int {
	f := util.NewScannerForInput(6, false)
	defer f.Close()

	product := 1

	races := ParseInput(f)
	for _, race := range races {
		product *= race.FindNumberOfPossibleWins()
	}

	return product
}

type Race struct {
	Time     int
	Distance int
}

func ParseInput(f *util.FileScanner) []Race {
	var races []Race

	for f.Scan() {
		curr := f.Text()
		fields := strings.Fields(curr)
		// if our list is empty, we are parsing the time
		if len(races) == 0 {
			for i := 1; i < len(fields); i++ {
				time, err := strconv.Atoi(fields[i])
				if err != nil {
					panic(err)
				}
				races = append(races, Race{
					Time: time,
				})
			}
		} else {
			// we are parsing the distance
			for i := 1; i < len(fields); i++ {
				distance, err := strconv.Atoi(fields[i])
				if err != nil {
					panic(err)
				}
				races[i-1].Distance = distance
			}
		}
	}
	return races
}

func (r Race) FindNumberOfPossibleWins() int {
	winCount := 0
	for timeHeldDown := 1; timeHeldDown < r.Time; timeHeldDown++ {
		if timeHeldDown*(r.Time-timeHeldDown) > r.Distance {
			winCount++
		}
	}
	return winCount
}
