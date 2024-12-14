package utils

import (
	"fmt"
	"strconv"
)

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			print(cell)
		}
		println()
	}
}

func CopyGrid(A [][]string) [][]string {
	B := make([][]string, len(A))
	for i := range A {
		B[i] = make([]string, len(A[i]))
		copy(B[i], A[i])
	}
	return B
}

type Coordinate [2]int

func Ints(A []string) []int {
	B := make([]int, len(A))
	for i, a := range A {
		B[i], _ = strconv.Atoi(a)
	}
	return B
}

func (a Coordinate) Sub(b Coordinate) Coordinate {
	return Coordinate{a[0] - b[0], a[1] - b[1]}
}

func (a Coordinate) Add(b Coordinate) Coordinate {
	return Coordinate{a[0] + b[0], a[1] + b[1]}
}
func (a Coordinate) Multiply(n int) Coordinate {
	return Coordinate{a[0] * n, a[1] * n}
}
func (a Coordinate) Equals(b Coordinate) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func (a Coordinate) InBounds(grid [][]string) bool {
	return a[0] >= 0 && a[0] < len(grid) && a[1] >= 0 && a[1] < len(grid[0])
}
func (a Coordinate) InBoundsInt(grid [][]int) bool {
	return a[0] >= 0 && a[0] < len(grid) && a[1] >= 0 && a[1] < len(grid[0])
}
func (a Coordinate) Unpack() (int, int) {
	return a[0], a[1]
}

func (a Coordinate) String() string {
	return fmt.Sprintf("%d,%d", a[0], a[1])
}
