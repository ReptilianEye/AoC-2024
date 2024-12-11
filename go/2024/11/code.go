package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func parseInput(input string) map[int]int {
	stonesStr := strings.Split(input, " ")
	stones := make(map[int]int)
	for _, s := range stonesStr {
		num, _ := strconv.Atoi(s)
		stones[num] += 1
	}
	return stones
}
func simulateRound(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for val, count := range stones {
		if val == 0 {
			newStones[1] += count
		} else if currStr := strconv.Itoa(val); len(currStr)%2 == 0 {
			left, _ := strconv.Atoi(currStr[:len(currStr)/2])
			right, _ := strconv.Atoi(currStr[len(currStr)/2:])
			newStones[left] += count
			newStones[right] += count
		} else {
			newStones[val*2024] += count
		}
	}
	return newStones
}

func solve(stones map[int]int, rounds int) int {
	for round := 1; round <= rounds; round++ {
		stones = simulateRound(stones)
	}
	cnt := 0
	for _, count := range stones {
		cnt += count
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
	stones := parseInput(input)
	if part2 {
		return solve(stones, 75)
	}
	// solve part 1 here
	return solve(stones, 25)
}
