package main

import (
	"regexp"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/thoas/go-funk"
)

func main() {
	aoc.Harness(run)
}

func solvePart1(input string) int {
	res := 0
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	for _, match := range r.FindAllStringSubmatch(input, -1) {
		v1, _ := strconv.Atoi(match[1])
		v2, _ := strconv.Atoi(match[2])
		res += v1 * v2
	}
	return res
}
func isDoCloser(i int, doS, dontS [][]int) bool {
	doDistance := i
	for _, do := range doS {
		if do[0] < i {
			doDistance = i - do[0]
		}
	}
	dontDistance := i
	for _, dont := range dontS {
		if dont[0] < i {
			dontDistance = i - dont[0]
		}
	}
	if dontDistance == -1 {
		return true
	}
	return doDistance <= dontDistance

}

func solvePart2(input string) int {
	res := 0
	rMul, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	rDo, _ := regexp.Compile(`do\(\)`)
	rDont, _ := regexp.Compile(`don't\(\)`)

	doS := rDo.FindAllStringIndex(input, -1)
	dontS := rDont.FindAllStringIndex(input, -1)

	matchesWithIndexes := funk.Zip(
		rMul.FindAllStringSubmatch(input, -1),
		rMul.FindAllStringIndex(input, -1),
	)

	for _, pair := range matchesWithIndexes {
		match := pair.Element1.([]string)
		index := pair.Element2.([]int)
		if isDoCloser(index[0], doS, dontS) {
			v1, _ := strconv.Atoi(match[1])
			v2, _ := strconv.Atoi(match[2])
			res += v1 * v2
		}
	}
	return res
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
		return solvePart2(input)
	}
	// solve part 1 here
	return solvePart1(input)
}
