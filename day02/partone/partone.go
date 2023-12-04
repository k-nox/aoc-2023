package partone

import (
	"fmt"
	"github.com/k-nox/aoc2023/util"
	"strconv"
	"strings"
)

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Solve() int {
	fscanner := util.NewFileScanner("day02/input/input.txt")
	defer fscanner.Close()

	sum := 0

	for fscanner.Scan() {
		curr := fscanner.Text()
		id := GetGameID(curr)
		validGame := true
		hands := GetHands(curr)
		for _, hand := range hands {
			if validGame == false {
				break
			}
			cubes := GetCubes(hand)
			// split each color/cube combo by space & convert first field to int - this is the number of that color
			for _, cube := range cubes {
				num, color := GetNumberOfCubesAndColor(cube)
				// use the map to determine if number is too much
				if maxCubes[color] < num {
					validGame = false
					break
				}
			}
		}
		if validGame {
			sum += id
		}
	}

	return sum
}

func GetGameID(line string) int {
	// split by colon to separate into game label and hands
	game := strings.Split(line, ":")
	// split label by space and convert second field to int - this is the game id
	label := strings.Fields(game[0])
	id, err := strconv.Atoi(label[1])
	if err != nil {
		fmt.Printf("game id: unable to convert %s to int: %s", label[1], err.Error())
		panic(err)
	}
	return id
}

func GetHands(line string) []string {
	game := strings.Split(line, ":")
	// split hands by semicolon - this represents each hand drawn
	return strings.Split(game[1], ";")
}

func GetCubes(hand string) []string {
	// split each hand by comma - this represents each individual color/cube combo
	return strings.Split(hand, ",")
}

func GetNumberOfCubesAndColor(cube string) (int, string) {
	fields := strings.Fields(cube)
	num, err := strconv.Atoi(fields[0])
	if err != nil {
		fmt.Printf("color cube combo: unable to convert %s to int: %s", fields[0], err.Error())
		panic(err)
	}
	return num, fields[1]
}
