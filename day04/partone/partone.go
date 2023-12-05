package partone

import (
	"github.com/k-nox/aoc2023/util"
	"math"
	"strings"
)

func Solve() int {
	points := 0

	fscanner := util.NewFileScanner("day04/input/input.txt")
	defer fscanner.Close()

	for fscanner.Scan() {
		n := 0
		card := fscanner.Text()
		cardNumbers := GetCardNumbers(card)
		winningNumbers := GetWinningNumbers(cardNumbers)
		playedNumbers := GetPlayedNumbers(cardNumbers)
		for _, played := range playedNumbers {
			if _, ok := winningNumbers[played]; ok {
				n++
			}
		}
		if n > 0 {
			points += int(math.Pow(2, float64(n-1)))
		}
	}
	return points
}

func GetCardNumbers(card string) string {
	return strings.Split(card, ":")[1]
}

func GetWinningNumbers(cardNumbers string) map[string]struct{} {
	numMap := map[string]struct{}{}
	winningSide := strings.Split(cardNumbers, "|")[0]
	nums := strings.Fields(winningSide)
	for _, num := range nums {
		numMap[num] = struct{}{}
	}
	return numMap
}

func GetPlayedNumbers(cardNumbers string) []string {
	playedSide := strings.Split(cardNumbers, "|")[1]
	return strings.Fields(playedSide)
}
