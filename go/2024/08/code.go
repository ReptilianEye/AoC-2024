package main

import (
	. "aoc-in-go/2024/utils"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}
func parseInput(input string) (A [][]string, antennas map[string][]Coordinate) {
	antennas = make(map[string][]Coordinate)

	lines := strings.Split(input, "\n")
	A = make([][]string, len(lines))
	for i, line := range lines {
		elements := strings.Split(line, "")
		A[i] = make([]string, len(elements))
		for j, element := range elements {
			A[i][j] = element
			if element != "." {
				antennas[element] = append(antennas[element], Coordinate{i, j})
			}
		}
	}
	return
}

func tryPlacingAntinodeV1(p1, p2 Coordinate, antinodes [][]string) {
	delta := p2.Sub(p1)
	newPos := p2.Add(delta)
	if newPos.InBounds(antinodes) {
		antinodes[newPos[0]][newPos[1]] = "#"
	}
}
func tryPlacingAntinodeV2(p1, p2 Coordinate, antinodes [][]string) {
	antinodes[p1[0]][p1[1]] = "#"
	antinodes[p2[0]][p2[1]] = "#"

	for {
		delta := p2.Sub(p1)
		newPos := p2.Add(delta)
		if !newPos.InBounds(antinodes) {
			return
		}
		x, y := newPos.Unpack()
		antinodes[x][y] = "#"
		p1 = p2
		p2 = Coordinate{x, y}
	}
}
func countAntinodes(antinodes [][]string) int {
	cnt := 0
	for _, row := range antinodes {
		for _, v := range row {
			if v == "#" {
				cnt += 1
			}
		}
	}
	return cnt
}
func solvePart1(A [][]string, anntenas map[string][]Coordinate) int {
	antinodes := CopyGrid(A)
	for _, positions := range anntenas {
		for i := 0; i < len(positions); i++ {
			for j := 0; j < len(positions); j++ {
				if i == j {
					continue
				}
				tryPlacingAntinodeV1(positions[i], positions[j], antinodes)
			}
		}
	}
	return countAntinodes(antinodes)
}
func solvePart2(A [][]string, anntenas map[string][]Coordinate) int {
	antinodes := CopyGrid(A)
	for _, positions := range anntenas {
		for i := 0; i < len(positions); i++ {
			for j := 0; j < len(positions); j++ {
				if i != j {
					tryPlacingAntinodeV2(positions[i], positions[j], antinodes)
				}
			}
		}
	}
	return countAntinodes(antinodes)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	A, anntenas := parseInput(input)
	if part2 {
		return solvePart2(A, anntenas)
	}
	// solve part 1 here
	return solvePart1(A, anntenas)
}
