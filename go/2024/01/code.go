package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/thoas/go-funk"
)

func main() {
	aoc.Harness(run)
}

func parseInput(input string) [][]int {
	in := [][]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		els := strings.Split(line, "   ")
		elsInt := funk.Map(els, func(el string) int {
			val, _ := strconv.Atoi(el)
			return val
		}).([]int)
		in = append(in, elsInt)
	}
	return in
}

func solvePart1(input [][]int) int {
	n := len(input)
	left := make([]int, n)
	right := make([]int, n)
	for i, pair := range input {
		left[i] = pair[0]
		right[i] = pair[1]
	}
	slices.Sort(left)
	slices.Sort(right)

	diff := 0
	for i := range left {
		diff += int(math.Abs(float64(left[i] - right[i])))
	}
	return diff
	// or using funk
	// return funk.Reduce(funk.Zip(left, right), func(acc int, pair []int) int {
	// 	return acc + funk.MaxInt(pair) - funk.MinInt(pair)
	// }, 0).(int)
}
func solvePart2(input [][]int) int {
	left := make([]int, len(input))
	counts := make(map[int]int)
	for i, pair := range input {
		left[i] = pair[0]
		counts[pair[1]]++
	}
	similarity := 0
	for _, l := range left {
		similarity += l * counts[l]
	}
	return similarity
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	parsedInput := parseInput(input)
	if part2 {
		return solvePart2(parsedInput)
	}
	// solve part 1 here
	return solvePart1(parsedInput)
}
