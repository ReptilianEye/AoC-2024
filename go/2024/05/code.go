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
func parseInput(input string) (graph map[int][]int, printsToCheck [][]int) {
	in := strings.Split(input, "\n\n")
	graph = map[int][]int{}
	for _, line := range strings.Split(in[0], "\n") {
		verts := strings.Split(line, "|")
		vertsInt := funk.Map(verts, func(v string) int {
			vInt, _ := strconv.Atoi(v)
			return vInt
		}).([]int)
		graph[vertsInt[0]] = append(graph[vertsInt[0]], vertsInt[1])
	}
	printsToCheck = [][]int{}
	for _, line := range strings.Split(in[1], "\n") {
		prints := strings.Split(line, ",")
		printsInt := funk.Map(prints, func(v string) int {
			vInt, _ := strconv.Atoi(v)
			return vInt
		}).([]int)
		printsToCheck = append(printsToCheck, printsInt)
	}
	return
}

func checkPrint(G map[int][]int, prints []int) bool {
	for i := len(prints) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if funk.Contains(G[prints[i]], prints[j]) {
				return false
			}
		}
	}
	return true
}
func solvePart1(G map[int][]int, printsToCheck [][]int) int {
	sol := 0
	for _, prints := range printsToCheck {
		if checkPrint(G, prints) {
			sol += prints[len(prints)/2]
		}
	}
	return sol
}

func topologicalSort(V []int, G map[int][]int) []int {
	visited := map[int]bool{}
	sortingStack := []int{}
	var topologicalSortUtil func(int)
	topologicalSortUtil = func(v int) {
		visited[v] = true

		for _, i := range G[v] {
			if !visited[i] && funk.Contains(V, i) {
				topologicalSortUtil(i)
			}
		}

		sortingStack = append([]int{v}, sortingStack...)
	}
	for _, v := range V {
		if !visited[v] {
			topologicalSortUtil(v)
		}
	}
	return sortingStack
}
func fixPrints(G map[int][]int, prints []int) []int {
	return topologicalSort(prints, G)
}
func solvePart2(G map[int][]int, printsToCheck [][]int) int {
	sol := 0
	for _, prints := range printsToCheck {
		if !checkPrint(G, prints) {
			fixed := fixPrints(G, prints)
			sol += fixed[len(fixed)/2]
		}
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
	// when you're ready to do part 2, remove this "not implemented" block
	graph, printsToCheck := parseInput(input)
	if part2 {
		return solvePart2(graph, printsToCheck)
	}
	// solve part 1 here
	return solvePart1(graph, printsToCheck)
}
