package parttwo

import (
	"github.com/k-nox/aoc2023/day05/partone"
	"github.com/k-nox/aoc2023/util"
	"math"
)

func Solve() int {
	f := util.NewFileScanner("day05/input/input.txt")
	defer f.Close()

	m, seedList := partone.ParseInput(f)
	seedRanges := parseSeedRanges(seedList)

	transformedRanges := Transform(m.Next, seedRanges)

	minLoc := math.MaxInt
	for _, r := range transformedRanges {
		if r.Start < minLoc {
			minLoc = r.Start
		}
	}

	return minLoc
}

type SeedRange struct {
	Start int
	Range int
}

func parseSeedRanges(seedList []int) []SeedRange {
	var seeds []SeedRange
	for i := 0; i < len(seedList)-1; i += 2 {
		seeds = append(seeds, SeedRange{
			Start: seedList[i],
			Range: seedList[i+1],
		})
	}
	return seeds
}

func Transform(m *partone.Map, srs []SeedRange) []SeedRange {
	var transformed []SeedRange

	// for each range, loop through each rule
	for _, sr := range srs {
		wasTransformed := false
		rangeEnd := sr.Start + sr.Range
		for _, r := range m.Rules {
			ruleSourceEnd := r.SourceStart + r.Range
			// possibilities:
			// 1. whole range is inside the rule -> transform the entire range & you're done looking at that range
			if r.SourceStart <= sr.Start && rangeEnd <= ruleSourceEnd {
				transformed = append(transformed, SeedRange{
					Start: r.DestStart + sr.Start - r.SourceStart,
					Range: sr.Range,
				})
				wasTransformed = true
				break
			}
			// 2. start is outside the rule, but middle to end is inside the rule -> end up with 2 seedranges, one that has not been transformed (start to middle) and one that has (middle to end)
			if sr.Start < r.SourceStart && r.SourceStart < rangeEnd && rangeEnd <= ruleSourceEnd {
				transformed = append(transformed,
					// add start -> the part of the range that's within the rule
					SeedRange{
						Start: sr.Start,
						Range: r.SourceStart - sr.Start,
					},
					// add part of the range within the rule -> end of range
					SeedRange{
						Start: r.DestStart,
						Range: rangeEnd - r.SourceStart,
					})
				wasTransformed = true
				break
			}
			// 3. start to middle is within the rule, but middle to end is not -> end up with 2 seedranges, one that has not been transformed (middle to end), and one that has (start to middle)
			if r.SourceStart <= sr.Start && sr.Start < ruleSourceEnd && ruleSourceEnd <= rangeEnd {
				transformed = append(transformed,
					// add start of range -> end of rule
					SeedRange{
						Start: r.DestStart + sr.Start - r.SourceStart,
						Range: ruleSourceEnd - sr.Start,
					},
					// add end of rule -> end of range
					SeedRange{
						Start: ruleSourceEnd,
						Range: rangeEnd - ruleSourceEnd,
					})
				wasTransformed = true
				break
			}
			// 4. start and end are outside rule, but middle is inside rule -> end up with three seedranges, 2 that have not been transformed (start to middle, middle to end), and one that has (middle)
			if sr.Start <= r.SourceStart && ruleSourceEnd <= rangeEnd {
				transformed = append(transformed,
					// range start -> rule start
					SeedRange{
						Start: sr.Start,
						Range: r.SourceStart - sr.Start,
					},
					// rule start -> rule end
					SeedRange{
						Start: r.DestStart,
						Range: r.Range,
					},
					// rule end -> range end
					SeedRange{
						Start: ruleSourceEnd,
						Range: rangeEnd - ruleSourceEnd,
					})
				wasTransformed = true
				break
			}
			// 5. range is entirely outside the rule, no action needed for this rule
		}
		if !wasTransformed {
			transformed = append(transformed, SeedRange{
				Start: sr.Start,
				Range: sr.Range,
			})
		}
	}

	if m.Next == nil {
		return transformed
	}

	return Transform(m.Next, transformed)
}
