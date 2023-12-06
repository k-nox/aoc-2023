package partone

import (
	"github.com/k-nox/aoc2023/util"
	"slices"
	"strconv"
	"strings"
)

func Solve() int {
	f := util.NewFileScanner("day05/input/input.txt")
	defer f.Close()

	m, seeds := ParseInput(f)
	locs := []int{}

	// loop over seeds
	for _, seed := range seeds {
		locs = append(locs, m.Transform(seed))
	}

	// now find the minimum location
	return slices.Min(locs)
}

type Rule struct {
	SourceStart int
	DestStart   int
	Range       int
}

type Map struct {
	Label string
	Next  *Map
	Rules []Rule
}

func ParseInput(f *util.FileScanner) (*Map, []int) {
	m := &Map{
		Label: "seed",
	}
	currMap := m
	seeds := []int{}
	for f.Scan() {
		curr := f.Text()
		// if we hit a blank line, create a new map, add it as the currMap's next, and set currMap to new map, then continue
		if curr == "" {
			next := &Map{}
			currMap.Next = next
			currMap = next
			continue
		}

		// special case to parse seeds
		if currMap.Label == "seed" {
			seedsArr := strings.Split(curr, ":")
			seedsStr := seedsArr[1]
			seedsStrArr := strings.Fields(seedsStr)
			for _, numStr := range seedsStrArr {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, num)
			}
			continue
		}

		// if our current map doesn't have a label, we are on a label line, so parse it
		if currMap.Label == "" {
			currMap.Label = GetLabel(curr)
			continue
		}

		// else we are inside the map, so parse the rule
		rule := ParseRule(curr)
		currMap.Rules = append(currMap.Rules, rule)
	}
	return m, seeds
}

func GetLabel(s string) string {
	return strings.Split(strings.Fields(s)[0], "-")[2]
}

func ParseRule(s string) Rule {
	nums := strings.Fields(s)
	dest, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}

	source, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	r, err := strconv.Atoi(nums[2])
	if err != nil {
		panic(err)
	}

	return Rule{
		SourceStart: source,
		DestStart:   dest,
		Range:       r,
	}
}

func (m *Map) Transform(num int) int {
	transformed := num
	for _, r := range m.Rules {
		// if the sourceStart <= seed < sourceStart+range, then seed = deststart + (seed - sourcestart), then go to next map
		if r.SourceStart <= num && num < r.SourceStart+r.Range {
			transformed = r.DestStart + (num - r.SourceStart)
			break
		}
	}
	// once there are no more maps, store the num we have as the location, and repeat the process
	if m.Next == nil {
		return transformed
	}
	return m.Next.Transform(transformed)
}
