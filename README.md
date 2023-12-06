# Advent of Code 2023

This repository holds my solutions in Go to the [Advent of Code 2023](https://adventofcode.com/) puzzles.

Each day follows a simple structure:
```
- day{num}
    - input
        - input.txt
        - sample.txt
    - partone
        - partone.go
    - parttwo
        - parttwo.go
    main.go
```

Each day is it's own small program with it's own `main` function and can be run independently with:
```shell
$ go run day<num>/main.go
```

## Completion Tracker
| Day | Part One Done | Part Two Done |
|-----|---------------|--------------|
| 1   | ✅             | ✅            |
| 2   | ✅             | ✅            |
| 3   | ✅             | ✅            |
| 4   | ✅             | ✅            |
| 5   | ✅             | ✅            |
| 6   | ✅             | ✅            |
| 7   |              |              |
| 8   |              |              |
| 9   |              |              |
| 10  |              |              |
| 11  |              |              |
| 12  |              |              |
| 13  |              |              |
| 14  |              |              |
| 15  |              |              |
| 16  |              |              |
| 17  |              |              |
| 18  |              |              |
| 19  |              |              |
| 20  |              |              |
| 21  |              |              |
| 22  |              |              |
| 23  |              |              |
| 24  |              |              |
| 25  |              |              |

## Day Generation
Starting with day 7, all the files in this repository were genereated with a simple generator using go text templates.
Checkout `util/gen.go` for the implementation.