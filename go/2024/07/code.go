package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/thoas/go-funk"
)

func main() {
	aoc.Harness(run)
}

type equation struct {
	result  int
	numbers []int
}

func parseInput(input string) []equation {
	equations := []equation{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		eq := equation{}
		split := strings.Split(line, ": ")
		eq.result, _ = strconv.Atoi(split[0])
		numbers := strings.Split(split[1], " ")
		eq.numbers = append(eq.numbers, funk.Map(numbers, func(v string) int {
			vInt, _ := strconv.Atoi(v)
			return vInt
		}).([]int)...)
		equations = append(equations, eq)
	}
	return equations
}

var operators = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}

func tryPlacingOperators(eq equation) bool {
	var placeOperator func(int, int) bool
	placeOperator = func(i, resultSoFar int) bool {
		if i == len(eq.numbers) {
			return resultSoFar == eq.result
		}
		for _, op := range operators {
			if placeOperator(i+1, op(resultSoFar, eq.numbers[i])) {
				return true
			}
		}
		return false
	}
	return placeOperator(1, eq.numbers[0])

}

func solvePart1(equations []equation) int {
	result := 0
	for _, eq := range equations {
		if tryPlacingOperators(eq) {
			result += eq.result
		}
	}
	return result
}
func solvePart2(equations []equation) int {
	operators["||"] = func(a, b int) int {
		aStr := strconv.Itoa(a)
		bStr := strconv.Itoa(b)
		r, _ := strconv.Atoi(aStr + bStr)
		return r
	}
	return solvePart1(equations)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	equations := parseInput(input)
	if part2 {
		return solvePart2(equations)
	}
	// solve part 1 here
	return solvePart1(equations)
}
