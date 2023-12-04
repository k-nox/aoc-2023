package parttwo

import (
	"github.com/k-nox/aoc2023/day02/partone"
	"github.com/k-nox/aoc2023/util"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

func Solve() int {
	fscanner := util.NewFileScanner("day02/input/input.txt")
	defer fscanner.Close()
	sum := 0

	for fscanner.Scan() {
		curr := fscanner.Text()
		counts := NewColorCountsMap()
		hands := partone.GetHands(curr)
		for _, hand := range hands {
			cubes := partone.GetCubes(hand)
			for _, cube := range cubes {
				// separate each num and color out
				num, color := partone.GetNumberOfCubesAndColor(cube)
				// if the num is greater than the current count, set a new count
				if counts[color] < num {
					counts[color] = num
				}
			}
		}
		// loop over the counts and get the power
		power := 1
		for _, count := range counts {
			power *= count
		}

		// add power to sum
		sum += power
	}

	return sum
}

func NewColorCountsMap() map[string]int {
	return map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}
}
