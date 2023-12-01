package main

import (
	"aoc2023/internal/inputfile"
	"fmt"
	"strings"
)

func main() {
	var sum = 0
	lines := strings.Split(inputfile.InputFile1, "\n")
	for _, line := range lines {
		var firstDigit = -1
		var lastDigit = -1
		for _, char := range line {
			if char >= '0' && char <= '9' {
				if firstDigit == -1 {
					firstDigit = int(char - '0')
				}
				lastDigit = int(char - '0')
			}
		}
		sum += firstDigit*10 + lastDigit
	}
	fmt.Printf("Part 1: %d\n", sum)
	letteredNumbers := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	sum = 0
	for _, line := range lines {
		var indexOfFirstDigit = -1
		var firstDigit = -1
		var indexOfLastDigit = -1
		var lastDigit = -1
		for number, letteredNumber := range letteredNumbers {
			index := strings.Index(line, letteredNumber)
			if index > -1 && (indexOfFirstDigit == -1 || index < indexOfFirstDigit) {
				indexOfFirstDigit = index
				firstDigit = number + 1
			}
			index = strings.LastIndex(line, letteredNumber)
			if index > -1 && (indexOfLastDigit == -1 || index > indexOfLastDigit) {
				indexOfLastDigit = index
				lastDigit = number + 1
			}
		}
		for index, char := range line {
			if char >= '0' && char <= '9' {
				if indexOfFirstDigit == -1 || index < indexOfFirstDigit {
					firstDigit = int(char - '0')
					indexOfFirstDigit = index
				}
				if indexOfLastDigit == -1 || index > indexOfLastDigit {
					lastDigit = int(char - '0')
					indexOfLastDigit = index
				}
			}
		}
		sum += firstDigit*10 + lastDigit
	}
	fmt.Printf("Part 2: %d\n", sum)
}
