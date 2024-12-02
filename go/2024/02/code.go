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

func solvePart1(reports [][]int) int {
	cnt := 0
	for _, report := range reports {
		if checkReport(report) {
			cnt++
		}
	}
	return cnt
}
func solvePart2(reports [][]int) int {
	cnt := 0
	for _, report := range reports {
		if checkReportDrop(report) {
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
		return solvePart2(reports)
		// return solvePart2Concurrent(reports)
	}
	// solve part 1 here
	return solvePart1(reports)
}
