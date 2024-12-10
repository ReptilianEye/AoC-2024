package main

import (
	. "aoc-in-go/2024/utils"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var steps = []Coordinate{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func parseInput(input string) (A [][]int, starts []Coordinate) {
	starts = []Coordinate{}

	lines := strings.Split(input, "\n")
	A = make([][]int, len(lines))
	for i, line := range lines {
		elements := strings.Split(line, "")
		A[i] = make([]int, len(elements))
		for j, element := range elements {
			A[i][j], _ = strconv.Atoi(element)
			if A[i][j] == 0 {
				starts = append(starts, Coordinate{i, j})
			}
		}
	}
	return
}

func countScoreForStart(s Coordinate, A [][]int) int {
	reachedTops := make(map[string]bool)
	var makeStep func(Coordinate)
	makeStep = func(current Coordinate) {
		cx, cy := current.Unpack()
		if A[current[0]][current[1]] == 9 {
			reachedTops[current.String()] = true
			return
		}
		for _, step := range steps {
			p := current.Add(step)
			if !p.InBoundsInt(A) {
				continue
			}
			x, y := p.Unpack()
			if A[x][y]-A[cx][cy] == 1 {
				makeStep(p)
			}
		}
	}
	makeStep(s)
	return len(reachedTops)

}
func countRatingForStart(s Coordinate, A [][]int) int {
	reachedTops := 0
	var makeStep func(Coordinate)
	makeStep = func(current Coordinate) {
		cx, cy := current.Unpack()
		if A[cx][cy] == 9 {
			reachedTops++
			return
		}
		for _, step := range steps {
			p := current.Add(step)
			if !p.InBoundsInt(A) {
				continue
			}
			x, y := p.Unpack()
			if A[x][y]-A[cx][cy] == 1 {
				makeStep(p)
			}
		}
	}
	makeStep(s)
	return reachedTops

}
func solvePart1(A [][]int, starts []Coordinate) int {
	cnt := 0
	for _, position := range starts {
		cnt += countScoreForStart(position, A)
	}
	return cnt
}
func solvePart2(A [][]int, starts []Coordinate) int {
	cnt := 0
	for _, position := range starts {
		cnt += countRatingForStart(position, A)
	}
	return cnt
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	A, starts := parseInput(input)
	if part2 {
		return solvePart2(A, starts)
	}
	// solve part 1 here
	return solvePart1(A, starts)
}
