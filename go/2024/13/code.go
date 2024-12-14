package main

import (
	. "aoc-in-go/2024/utils"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}
func parseInput(input string) [][3]Coordinate {
	lines := strings.Split(input, "\n\n")
	parsedInput := [][3]Coordinate{}
	for _, line := range lines {
		buttonRegex := regexp.MustCompile(`Button [A-Z]: X\+([0-9]+), Y\+([0-9]+)`)
		prizeRegex := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)

		buttons := []Coordinate{}
		matches := buttonRegex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			button := Coordinate{x, y}
			buttons = append(buttons, button)
		}

		prizeMatch := prizeRegex.FindStringSubmatch(line)
		prizeX, _ := strconv.Atoi(prizeMatch[1])
		prizeY, _ := strconv.Atoi(prizeMatch[2])
		prize := Coordinate{prizeX, prizeY}
		parsedInput = append(parsedInput, [3]Coordinate{buttons[0], buttons[1], prize})
	}
	return parsedInput
}
func bestSolution(query [3]Coordinate) (int, bool) {
	a, b, prize := query[0], query[1], query[2]
	bestSol := -1
	found := false
	for aClicks := 0; aClicks <= 100; aClicks++ {
		for bClicks := 0; bClicks <= 100; bClicks++ {
			aSteps := a.Multiply(aClicks)
			bSteps := b.Multiply(bClicks)
			final := aSteps.Add(bSteps)
			if final.Equals(prize) {
				currentPrice := 3*aClicks + bClicks
				if !found || bestSol > currentPrice {
					bestSol = currentPrice
					found = true
				}
			}
		}
	}
	return bestSol, found

}
func solvePart1(query [][3]Coordinate) int {
	totalTokens := 0
	for _, q := range query {
		best, success := bestSolution(q)
		if success {
			totalTokens += best
		}
	}
	return totalTokens
}

type point [2]float64

func findIntersection(p1, p2, p3 Coordinate) (point, bool) {
	a0, a1 := float64(p1[0]), float64(p1[1])
	b0, b1 := float64(p2[0]), float64(p2[1])
	c0, c1 := float64(p3[0]), float64(p3[1])

	det := a0*b1 - a1*b0

	if det == 0.0 {
		return point{0.0, 0.0}, false
	}

	return point{(c0*b1 - c1*b0) / det, (a0*c1 - a1*c0) / det}, true

}

func isCloseToInt(f float64) bool {
	return f == float64(int(f))
}

func solvePart2(query [][3]Coordinate) int {
	totalTokens := 0
	for _, q := range query {
		q[2] = q[2].Add(Coordinate{10000000000000, 10000000000000})
		intersectF, found := findIntersection(q[0], q[1], q[2])
		if found && isCloseToInt(intersectF[0]) && isCloseToInt(intersectF[1]) {
			aClicks := int(intersectF[0])
			bClicks := int(intersectF[1])
			totalTokens += 3*aClicks + bClicks
		}
	}
	return totalTokens
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	query := parseInput(input)
	if part2 {
		return solvePart2(query)
	}
	// solve part 1 here
	return solvePart1(query)
}
