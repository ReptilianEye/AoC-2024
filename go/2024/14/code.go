package main

import (
	. "aoc-in-go/2024/utils"
	"fmt"
	"regexp"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}
func parseInput(input string) [][2]Coordinate {
	lines := strings.Split(input, "\n")
	parsedInput := [][2]Coordinate{}
	for _, line := range lines {
		lineRegex := regexp.MustCompile(`p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)`)
		matches := lineRegex.FindAllStringSubmatch(line, -1)
		matchesInt := Ints(matches[0])
		startX, startY, velX, velY := matchesInt[0], matchesInt[1], matchesInt[2], matchesInt[3]

		parsedInput = append(
			parsedInput,
			[2]Coordinate{{startX, startY}, {velX, velY}},
		)
	}
	return parsedInput
}
func solvePart1(n, m, seconds int, robots [][2]Coordinate, print bool) int {
	mapOfRobots := make(map[Coordinate][]int)
	for i, robot := range robots {
		mapOfRobots[robot[0]] = append(mapOfRobots[robot[0]], i)
	}
	fmt.Println(mapOfRobots)
	for i := 1; i <= seconds; i++ {
		newMapOfRobots := make(map[Coordinate][]int)
		for pos, robotsOnPos := range mapOfRobots {
			for _, robotIdx := range robotsOnPos {
				robotMove := robots[robotIdx][1]
				newPos := pos.Add(robotMove)
				newPos = Coordinate{(newPos[0] + n) % n, (newPos[1] + m) % m}
				newMapOfRobots[newPos] = append(newMapOfRobots[newPos], robotIdx)
			}
		}
		mapOfRobots = newMapOfRobots
		if print {
			fmt.Println("after", i, "seconds")
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					c := Coordinate{j, i}
					if _, ok := newMapOfRobots[c]; ok {
						fmt.Print("#")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
			fmt.Println()
			fmt.Println()
		}

	}
	quadrants := [2][2]int{{0, 0}, {0, 0}}
	for pos, robotsOnPos := range mapOfRobots {
		if pos[0] < n/2 {
			if pos[1] < m/2 {
				quadrants[0][0] += len(robotsOnPos)
			} else if pos[1] > m/2 {
				quadrants[0][1] += len(robotsOnPos)
			}
		} else if pos[0] > n/2 {
			if pos[1] < m/2 {
				quadrants[1][0] += len(robotsOnPos)
			} else if pos[1] > m/2 {
				quadrants[1][1] += len(robotsOnPos)
			}
		}
	}
	sol := 1
	for _, quadrant := range quadrants {
		sol *= quadrant[0] * quadrant[1]
	}
	return sol
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	robots := parseInput(input)
	if part2 {
		return solvePart1(101, 103, 10000, robots, true)
	}
	return solvePart1(101, 103, 100, robots, false)
}
