package main

import (
	"aoc2023/internal/inputfile"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(strings.ReplaceAll(inputfile.InputFile3, "\r", ""), "\n")
	day3part1(lines)
	day3part2(lines)
}

func day3part1(lines []string) {
	sum := 0
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		partNumberStart := -1
		for x := 0; x < len(line); x++ {
			if line[x] < '0' || line[x] > '9' {
				if partNumberStart == -1 {
					continue
				}
				if checkSurroundings(lines, partNumberStart, x-1, y) {
					parsedNumber, _ := strconv.Atoi(line[partNumberStart:x])
					sum += parsedNumber
				}
				partNumberStart = -1
				continue
			}

			if partNumberStart == -1 {
				partNumberStart = x
			}
		}

		if partNumberStart != -1 && checkSurroundings(lines, partNumberStart, len(line)-1, y) {
			parsedNumber, _ := strconv.Atoi(line[partNumberStart:])
			sum += parsedNumber
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func checkSurroundings(lines []string, x1 int, x2 int, y int) bool {
	for _y := max(y-1, 0); _y <= min(y+1, len(lines)-1); _y++ {
		line := lines[_y]
		if _y == y {
			if (x1-1 >= 0 && isSymbol(line[x1-1])) ||
				(x2+1 < len(line) && isSymbol(line[x2+1])) {
				return true
			}
			continue
		}

		for x := max(x1-1, 0); x <= min(x2+1, len(line)-1); x++ {
			c := line[x]
			if isSymbol(c) {
				return true
			}
		}
	}
	return false
}

func isSymbol(c byte) bool {
	return c != '.' && (c < '0' || c > '9')
}

func day3part2(lines []string) {
	sum := 0
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			sum += getGearRatio(lines, x, y)
		}
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func findPartNumber(line string, x int) (int, int) {
	if line[x] < '0' || line[x] > '9' {
		return -1, -1
	}

	startX := x
	for ; startX >= 0; startX-- {
		if line[startX] < '0' || line[startX] > '9' {
			break
		}
	}
	startX++

	endX := x
	for ; endX < len(line); endX++ {
		if line[endX] < '0' || line[endX] > '9' {
			break
		}
	}
	endX--

	return startX, endX
}

func getGearRatio(lines []string, x int, y int) int {
	if lines[y][x] != '*' {
		return 0
	}
	firstRatio := -1
	secondRatio := -1

	for _y := max(y-1, 0); _y <= min(y+1, len(lines)-1); _y++ {
		line := lines[_y]
		for _x := max(x-1, 0); _x <= min(x+1, len(line)-1); _x++ {
			if _x == x && _y == y {
				continue
			}

			pnStart, pnEnd := findPartNumber(line, _x)
			if pnStart == -1 {
				continue
			}
			_x = pnEnd
			ratio, _ := strconv.Atoi(line[pnStart : pnEnd+1])
			if firstRatio == -1 {
				firstRatio = ratio
			} else if secondRatio == -1 {
				secondRatio = ratio
			} else {
				// More than two ratios, abort
				return 0
			}
		}
	}

	if secondRatio == -1 {
		// Not enough ratios found
		return 0
	}
	return firstRatio * secondRatio
}
