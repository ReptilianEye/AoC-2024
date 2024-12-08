package main

import (
	"aoc-in-go/2024/utils"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func parseInput(input string) (si, sj int, A [][]string) {
	lines := strings.Split(input, "\n")
	A = make([][]string, len(lines))
	for i, line := range lines {
		elements := strings.Split(line, "")
		A[i] = make([]string, len(elements))
		for j, element := range elements {
			A[i][j] = element
			if element == "^" {
				si, sj = i, j
			}
		}
	}
	return
}

var steps = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var stepsStr = [4]string{"^", ">", "v", "<"}

func makeStep(currentPos, step [2]int) [2]int {
	return [2]int{currentPos[0] + step[0], currentPos[1] + step[1]}
}
func simulateGuard(si, sj int, A [][]string) {
	isSaveIdx := func(i, j int) bool {
		return i >= 0 && i < len(A) && j >= 0 && j < len(A[i])
	}
	stepIdx := 0
	prev := [2]int{si, sj}
	for {
		A[prev[0]][prev[1]] = "X"
		nextStep := makeStep(prev, steps[stepIdx])
		if !isSaveIdx(nextStep[0], nextStep[1]) {
			return
		}
		if A[nextStep[0]][nextStep[1]] == "#" {
			stepIdx = (stepIdx + 1) % 4
			continue
		}
		prev = [2]int{nextStep[0], nextStep[1]}
	}
}

func checkIfGuardCycles(si, sj int, A [][]string) bool {
	isSaveIdx := func(i, j int) bool {
		return i >= 0 && i < len(A) && j >= 0 && j < len(A[i])
	}
	stepIdx := 0
	prev := [2]int{si, sj}
	for {
		nextStep := makeStep(prev, steps[stepIdx])
		if !isSaveIdx(nextStep[0], nextStep[1]) {
			return false
		}
		for A[nextStep[0]][nextStep[1]] == "#" {
			stepIdx = (stepIdx + 1) % 4
			nextStep = makeStep(prev, steps[stepIdx])
		}
		prev = [2]int{nextStep[0], nextStep[1]}
		if A[prev[0]][prev[1]] == stepsStr[stepIdx] {
			return true
		}
		A[prev[0]][prev[1]] = stepsStr[stepIdx]
	}
}

func solvePart1(si, sj int, A [][]string) int {
	simulateGuard(si, sj, A)
	count := 0
	for _, row := range A {
		for _, cell := range row {
			if cell == "X" {
				count++
			}
		}
	}
	return count
}

func solvePart2(si, sj int, A [][]string) int {
	count := 0
	for i := range A {
		for j := range A[i] {
			if A[i][j] == "." {
				B := utils.CopyGrid(A)
				B[i][j] = "#"
				if checkIfGuardCycles(si, sj, B) {
					count++
				}
			}
		}
	}
	return count
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return solvePart2(parseInput(input))
	}
	// solve part 1 here
	return solvePart1(parseInput(input))
}
