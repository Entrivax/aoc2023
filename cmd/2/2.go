package main

import (
	"aoc2023/internal/inputfile"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(inputfile.InputFile2, "\n")
	day2part1(lines)
	day2part2(lines)
}

func day2part1(lines []string) {
	sum := 0
	maxBlue := 14
	maxGreen := 13
	maxRed := 12
GameLoop:
	for _, line := range lines {
		gameId, _ := strconv.Atoi(line[5:strings.IndexRune(line, ':')])
		rounds := strings.Split(line[strings.IndexRune(line, ':')+1:], ";")

		for _, round := range rounds {
			for _, color := range strings.Split(round, ",") {
				trimedColor := strings.TrimSpace(color)
				count, _ := strconv.Atoi(trimedColor[:strings.IndexRune(trimedColor, ' ')])
				if strings.Contains(color, "blue") {
					if count > maxBlue {
						continue GameLoop
					}
				} else if strings.Contains(color, "red") {
					if count > maxRed {
						continue GameLoop
					}
				} else if strings.Contains(color, "green") {
					if count > maxGreen {
						continue GameLoop
					}
				}
			}
		}
		sum += gameId
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func day2part2(lines []string) {
	sum := 0
	for _, line := range lines {
		rounds := strings.Split(line[strings.IndexRune(line, ':')+1:], ";")
		minBlue := 0
		minGreen := 0
		minRed := 0

		for _, round := range rounds {
			for _, color := range strings.Split(round, ",") {
				trimedColor := strings.TrimSpace(color)
				count, _ := strconv.Atoi(trimedColor[:strings.IndexRune(trimedColor, ' ')])
				if strings.Contains(color, "blue") {
					if count > minBlue {
						minBlue = count
					}
				} else if strings.Contains(color, "red") {
					if count > minRed {
						minRed = count
					}
				} else if strings.Contains(color, "green") {
					if count > minGreen {
						minGreen = count
					}
				}
			}
		}
		sum += minBlue * minGreen * minRed
	}

	fmt.Printf("Part 2: %d\n", sum)
}
