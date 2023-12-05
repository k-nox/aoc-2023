package parttwo

import (
	"github.com/k-nox/aoc2023/day04/partone"
	"github.com/k-nox/aoc2023/util"
	"strconv"
	"strings"
)

func Solve() int {
	cardsWon := 0

	fs := util.NewFileScanner("day04/input/input.txt")
	defer fs.Close()

	copyCounts := map[int]int{}

	for fs.Scan() {
		matches := 0
		card := fs.Text()
		splitCard := strings.Split(card, ":")
		cardId := getID(splitCard[0])
		cardNumbers := splitCard[1]
		if _, ok := copyCounts[cardId]; !ok {
			// we have not seen this id yet
			copyCounts[cardId] = 0
		}
		copyCounts[cardId]++
		winningNumbers := partone.GetWinningNumbers(cardNumbers)
		playedNumbers := partone.GetPlayedNumbers(cardNumbers)
		for _, played := range playedNumbers {
			if _, ok := winningNumbers[played]; ok {
				matches += 1
			}
		}

		for i := 1; i <= matches; i++ {
			if _, ok := copyCounts[cardId+i]; !ok {
				copyCounts[cardId+i] = 0
			}
			copyCounts[cardId+i] += copyCounts[cardId]
		}
	}

	for _, count := range copyCounts {
		cardsWon += count
	}

	return cardsWon
}

func getID(label string) int {
	idStr := strings.Fields(label)[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	return id
}
