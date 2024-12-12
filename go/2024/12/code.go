package main

import (
	. "aoc-in-go/2024/utils"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var steps = []Coordinate{
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 0},
}

func parseInput(input string) [][]string {

	lines := strings.Split(input, "\n")
	A := make([][]string, len(lines))
	for i, line := range lines {
		elements := strings.Split(line, "")
		A[i] = make([]string, len(elements))
		for j, element := range elements {
			A[i][j] = element
		}
	}
	return A
}

func calculatePeriAndArea(s Coordinate, A [][]string) (perimeter, area int) {
	currentMarker := A[s[0]][s[1]]
	visited := make(map[Coordinate]bool)
	var makeStep func(Coordinate)
	makeStep = func(current Coordinate) {
		cx, cy := current.Unpack()
		area++
		A[cx][cy] = "."
		visited[current] = true
		for _, step := range steps {
			p := current.Add(step)
			if !p.InBounds(A) {
				perimeter++
				continue
			}
			x, y := p.Unpack()
			if A[x][y] == currentMarker {
				makeStep(p)
			} else {
				if _, ok := visited[p]; !ok {
					perimeter++
				}
			}

		}
	}
	makeStep(s)
	return perimeter, area

}
func solvePart1(A [][]string) int {
	result := 0
	for i := range A {
		for j := range A[i] {
			if A[i][j] != "." {
				peri, area := calculatePeriAndArea(Coordinate{i, j}, A)
				result += area * peri
			}
		}
	}
	return result
}
func countVertices(A [][]string, p Coordinate) int {

	marker := A[p[0]][p[1]]
	cnt := 0

	for i := 0; i < 4; i++ {
		toCheck := []Coordinate{p.Add(steps[i]), p.Add(steps[(i+1)%4])}

		// a node is a vertex if it has 2 neighbors that are not from the same marker
		// O O O
		// O X X - middle X is a outer vertex
		// O X O
		passed := true
		for _, c := range toCheck {
			if c.InBounds(A) && A[c[0]][c[1]] == marker {
				passed = false
			}
		}
		if passed {
			cnt++ // found a outer vertex
		}

		// a node is a vertex if it has 2 neighbors that are the same marker but the diagonal is not
		// X X - top left X is a inner vertex
		// X O
		passed = true
		for _, c := range toCheck {
			if c.InBounds(A) && A[c[0]][c[1]] == marker {
				continue
			}
			passed = false
		}
		if passed {
			diagonalStep := p.Add(steps[i]).Add(steps[(i+1)%4])
			if diagonalStep.InBounds(A) && A[diagonalStep[0]][diagonalStep[1]] != marker {
				cnt++ // found a vertex
			}
		}
	}
	return cnt
}
func calculateAreaAndSides(s Coordinate, A [][]string) (sides, area int) {
	currentMarker := A[s[0]][s[1]]
	B := CopyGrid(A)
	visited := make(map[Coordinate]bool)

	var makeStep func(Coordinate)
	makeStep = func(current Coordinate) {
		cx, cy := current.Unpack()
		A[cx][cy] = "."
		visited[current] = true
		for _, step := range steps {
			p := current.Add(step)
			if !p.InBounds(A) {
				continue
			}
			x, y := p.Unpack()
			if A[x][y] == currentMarker {
				makeStep(p)
			}
		}
	}

	makeStep(s)
	for p := range visited {
		s := countVertices(B, p) // number of sides == number of vertices
		sides += s
	}
	area = len(visited)
	return sides, area
}

func solvePart2(A [][]string) int {
	result := 0
	for i := range A {
		for j := range A[i] {
			if A[i][j] != "." {
				sides, area := calculateAreaAndSides(Coordinate{i, j}, A)
				result += area * sides
			}
		}
	}
	return result
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
