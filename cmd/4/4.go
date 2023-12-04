package main

import (
	"aoc2023/internal/inputfile"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(strings.ReplaceAll(inputfile.InputFile4, "\r", ""), "\n")
	day4part1(lines)
	day4part2(lines)
}

func day4part1(lines []string) {
	sum := 0
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		cardData := strings.TrimSpace(strings.Split(line, ": ")[1])
		cardNumbers := strings.Split(cardData, " | ")
		winningNumbers := parseSpacedNumbers(cardNumbers[0])
		ourNumbers := parseSpacedNumbers(cardNumbers[1])
		roundValue := 0
		for _, number := range ourNumbers {
			if slices.Contains(winningNumbers, number) {
				if roundValue == 0 {
					roundValue = 1
				} else {
					roundValue *= 2
				}
			}
		}
		sum += roundValue
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func parseSpacedNumbers(numbers string) []int {
	re := regexp.MustCompile(`\s+`)
	splitted := re.Split(numbers, -1)
	numbersArr := make([]int, len(splitted))
	for index, numberStr := range splitted {
		parsed, _ := strconv.Atoi(strings.TrimSpace(numberStr))
		numbersArr[index] = parsed
	}
	return numbersArr
}

func day4part2(lines []string) {
	processCount := make([]int, len(lines))
	for y := 0; y < len(lines); y++ {
		processCount[y] = 1
	}

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		cardData := strings.TrimSpace(strings.Split(line, ": ")[1])
		cardNumbers := strings.Split(cardData, " | ")
		winningNumbers := parseSpacedNumbers(cardNumbers[0])
		ourNumbers := parseSpacedNumbers(cardNumbers[1])
		wonCopies := 0
		for _, number := range ourNumbers {
			if slices.Contains(winningNumbers, number) {
				wonCopies++
			}
		}
		for i := 0; i < wonCopies; i++ {
			processCount[y+i+1] += processCount[y]
		}
	}

	sum := 0
	for _, count := range processCount {
		sum += count
	}

	fmt.Printf("Part 2: %d\n", sum)
}
