package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/thoas/go-funk"
)

func main() {
	aoc.Harness(run)
}
func parseInput(input string) [][]int {
	report := [][]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		levels := strings.Split(line, " ")
		levelsInt := funk.Map(levels, func(el string) int {
			val, _ := strconv.Atoi(el)
			return val
		}).([]int)
		report = append(report, levelsInt)
	}
	return report
}

func checkReport(report []int) bool {
	if !slices.IsSorted(report) && !slices.IsSorted(funk.ReverseInt(report)) {
		return false
	}
	prev := report[0]
	for i := 1; i < len(report); i++ {
		diff := int(math.Abs(float64(prev - report[i])))
		if diff < 1 || diff > 3 {
			return false
		}
		prev = report[i]
	}
	return true
}
func checkReportDrop(report []int) bool {
	for i := 0; i < len(report); i++ {
		current := slices.Clone(report)
		current = append(current[:i], current[i+1:]...)
		if checkReport(current) {
			return true
		}
	}
	return false
}

func checkReportDP(report []int) bool {
	return isSortedWithDrop(report) || isSortedWithDrop(funk.Reverse(report).([]int))
}

func isSortedWithDrop(report []int) bool {
	checkDiff := func(diff int) bool {
		return diff >= 1 && diff <= 3
	}
	n := len(report)

	DP := make([][]bool, n)
	for i := 0; i < n; i++ {
		DP[i] = make([]bool, 2)
	}
	DP[n-1][0] = true
	DP[n-1][1] = true

	DP[n-2][0] = checkDiff(report[n-1] - report[n-2])
	DP[n-2][1] = true

	for i := n - 3; i >= 0; i-- {
		canJump1 := checkDiff(report[i+1] - report[i])
		canJump2 := checkDiff(report[i+2] - report[i])
		if canJump1 {
			DP[i][0] = DP[i+1][0]
		} else {
			DP[i][0] = false
		}
		if canJump2 {
			DP[i][1] = DP[i+2][0]
		} else {
			DP[i][1] = false
		}
	}
	for i := 0; i < n-1; i++ {
		if !checkDiff(report[i+1] - report[i]) {
			res := DP[i][1]
			if i == 0 {
				return res || DP[1][0]
			} else {
				return res || DP[i-1][1]
			}
		}
	}
	return true
}

func solvePart1(reports [][]int) int {
	cnt := 0
	for _, report := range reports {
		if checkReport(report) {
			cnt++
		}
	}
	return cnt
}
func solvePart2(reports [][]int, useChecker ...func([]int) bool) int {
	checker := checkReportDrop
	if len(useChecker) > 0 {
		checker = useChecker[0]
	}
	cnt := 0
	for _, report := range reports {
		if checker(report) {
			cnt++
		}
	}
	return cnt
}
func worker(jobs <-chan []int, results chan<- bool, wg *sync.WaitGroup) {
	for report := range jobs {
		results <- checkReportDrop(report)
	}
	wg.Done()
}
func solvePart2Concurrent(reports [][]int) int {
	cnt := 0
	jobs := make(chan []int, len(reports))
	results := make(chan bool, len(reports))
	wg := sync.WaitGroup{}
	for _, report := range reports {
		jobs <- report
	}
	close(jobs)
	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}
	wg.Wait()
	close(results)
	for res := range results {
		if res {
			cnt++
		}
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
	reports := parseInput(input)
	// when you're ready to do part 2, remove this "not implemented" block

	if part2 {
		return solvePart2(reports, checkReportDP)
		// return solvePart2Concurrent(reports)
	}
	// solve part 1 here
	return solvePart1(reports)
}
