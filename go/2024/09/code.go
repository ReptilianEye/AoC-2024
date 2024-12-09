package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Range struct {
	start int
	end   int
	value string
}

func (r1 Range) len() int {
	return r1.end - r1.start + 1
}
func (r1 Range) fillRange(r2 Range) (result []Range, filled bool, copyLeftover Range) {
	if r1.len() >= r2.len() {
		leftPart := Range{
			start: r1.start,
			end:   r1.start + r2.len() - 1,
			value: r2.value,
		}
		leftoverSpace := Range{
			start: leftPart.end + 1,
			end:   r1.end,
			value: r1.value,
		}
		result = append(result, leftPart)
		if leftoverSpace.len() > 0 {
			result = append(result, leftoverSpace)
		}
		return result, true, Range{}
	} else {
		r1.value = r2.value
		return []Range{r1}, false, Range{
			start: r2.start,
			end:   r2.end - r1.len(),
			value: r2.value,
		}
	}
}

func parseInput(input string) []Range {
	values := strings.Split(input, "")
	id := 0
	isFile := true
	idx := 0
	ranges := []Range{}
	for _, v := range values {
		space, _ := strconv.Atoi(v)
		r := Range{
			start: idx,
			end:   idx + space - 1,
		}
		if isFile {
			r.value = strconv.Itoa(id)
			isFile = false
			id++
		} else {
			r.value = "."
			isFile = true
		}
		if r.len() > 0 {
			ranges = append(ranges, r)
		}
		idx += space
	}
	return ranges
}
func parseInput2(input string) ([]string, []Range) {
	values := strings.Split(input, "")
	id, idx, isFile := 0, 0, true
	files := []Range{}
	magistral := []string{}
	for _, v := range values {
		space, _ := strconv.Atoi(v)
		r := Range{
			start: idx,
			end:   idx + space - 1,
		}
		if isFile {
			r.value = strconv.Itoa(id)
			id++
		} else {
			r.value = "."
		}
		isFile = !isFile
		if r.len() > 0 {
			if r.value != "." {
				files = append(files, r)
			}
			for i := 0; i < r.len(); i++ {
				magistral = append(magistral, r.value)
			}

		}
		idx += space
	}
	return magistral, files
}
func solvePart1(ranges []Range) int {
	findNextNotEmpty := func() int {
		for i := len(ranges) - 1; i >= 0; i-- {
			if ranges[i].value != "." {
				return i
			}
		}
		return -1
	}
	for i := 0; i < len(ranges); i++ {
		lastNotEmpty := findNextNotEmpty()
		if lastNotEmpty == -1 {
			break
		}
		if ranges[i].value == "." {
			if i == lastNotEmpty-1 {
				r := Range{
					start: ranges[i].start,
					end:   ranges[i].start + ranges[lastNotEmpty].len() - 1,
					value: ranges[lastNotEmpty].value,
				}
				ranges = ranges[:i]
				ranges = append(ranges, r)
				break
			}
			fillingResult, filled, leftover := ranges[i].fillRange(ranges[lastNotEmpty])
			if filled {
				ranges = append(
					ranges[:lastNotEmpty],
					ranges[lastNotEmpty+1:]...) // remove lastNotEmpty - it was copied completely
				ranges = append(
					ranges[:i],
					append(fillingResult, ranges[i+1:]...)...) // insert copied range
			} else {
				ranges[i] = fillingResult[0]     // replace current range with copied range (only part of it)
				ranges[len(ranges)-1] = leftover // replace last range with leftover
			}
		}
		for ranges[len(ranges)-1].value == "." { // remove all empty ranges at the end
			ranges = ranges[:len(ranges)-1]
		}
	}
	answer := 0
	for i := 0; i < len(ranges); i++ {
		if ranges[i].value == "." {
			continue
		}
		for idx := ranges[i].start; idx <= ranges[i].end; idx++ {
			v, _ := strconv.Atoi(ranges[i].value)
			answer += idx * v
		}
	}

	return answer
}

func solvePart2(magistral []string, files []Range) int {
	slotSize := func(i int) int {
		size := 0
		for j := i; j < len(magistral); j++ {
			if magistral[j] != "." {
				break
			}
			size++
		}
		return size
	}
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		for j := 0; j < len(magistral) && j < file.start; j++ {
			if magistral[j] == "." {
				size := slotSize(j)
				if size >= file.len() {
					for k := 0; k < file.len(); k++ {
						magistral[j+k] = file.value
						magistral[file.start+k] = "."
					}
					break
				} else {
					j += size
				}
			}

		}
	}
	answer := 0
	for i := 0; i < len(magistral); i++ {
		if magistral[i] == "." {
			continue
		}
		v, _ := strconv.Atoi(string(magistral[i]))
		answer += i * v
	}

	return answer
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
		return solvePart2(parseInput2(input))
	}
	// solve part 1 here
	return solvePart1(parseInput(input))
}
