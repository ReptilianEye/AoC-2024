package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/thoas/go-funk"
)

func main() {
	aoc.Harness(run)
}

func countXMAS(s string) int {
	return strings.Count(s, "XMAS")
}
func createLine(A [][]string, currIdx []int, step []int) string {
	isSaveIdx := func(i, j int) bool {
		return i >= 0 && i < len(A) && j >= 0 && j < len(A[i])
	}
	s := ""
	for isSaveIdx(currIdx[0], currIdx[1]) {
		s += A[currIdx[0]][currIdx[1]]
		currIdx[0] += step[0]
		currIdx[1] += step[1]
	}
	return s

}

func solvePart1(A [][]string) int {
	cnt := 0
	n := len(A)
	for i := 0; i < len(A); i++ {
		toCheck := []string{}
		toCheck = append(toCheck, createLine(A, []int{i, 0}, []int{0, 1}))
		toCheck = append(toCheck, createLine(A, []int{0, i}, []int{1, 0}))
		toCheck = append(toCheck, createLine(A, []int{i, 0}, []int{1, 1}))
		if i != 0 {
			toCheck = append(toCheck, createLine(A, []int{0, i}, []int{1, 1}))
		}
		toCheck = append(toCheck, createLine(A, []int{0, n - i - 1}, []int{1, -1}))
		if i != 0 {
			toCheck = append(toCheck, createLine(A, []int{i, n - 1}, []int{1, -1}))
		}

		for _, s := range toCheck {
			cnt += countXMAS(s)
			cnt += countXMAS(funk.ReverseString(s))
		}
	}
	return cnt
}

func solvePart2(A [][]string) int {
	isSaveIdx := func(i, j int) bool {
		return i >= 0 && i < len(A) && j >= 0 && j < len(A[i])
	}
	checkCross := func(i, j int, steps [][]int) bool {
		indices := funk.Map(steps, func(step []int) []int {
			return []int{i + step[0], j + step[1]}
		}).([][]int)

		for _, idx := range indices {
			if !isSaveIdx(idx[0], idx[1]) {
				return false
			}
		}
		cross1 := A[indices[0][0]][indices[0][1]] + A[indices[2][0]][indices[2][1]]
		cross2 := A[indices[1][0]][indices[1][1]] + A[indices[3][0]][indices[3][1]]
		return (cross1 == "MS" || cross1 == "SM") && (cross2 == "MS" || cross2 == "SM")
	}

	steps := [][]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}} // diagonal
	cnt := 0
	for i := 1; i < len(A)-1; i++ {
		for j := 1; j < len(A[i])-1; j++ {
			if A[i][j] == "A" {
				if checkCross(i, j, steps) {
					cnt++
				}
			}
		}
	}
	return cnt

}
func parseInput(input string) [][]string {
	return funk.Map(strings.Split(input, "\n"), func(s string) []string {
		return strings.Split(s, "")
	}).([][]string)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	A := parseInput(input)
	if part2 {
		return solvePart2(A)
	}
	// solve part 1 here
	return solvePart1(A)
}
