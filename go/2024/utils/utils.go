package utils

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

func (a Coordinate) Sub(b Coordinate) Coordinate {
	return Coordinate{a[0] - b[0], a[1] - b[1]}
}

func (a Coordinate) Add(b Coordinate) Coordinate {
	return Coordinate{a[0] + b[0], a[1] + b[1]}
}

func (a Coordinate) Equals(b Coordinate) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func (a Coordinate) InBounds(grid [][]string) bool {
	return a[0] >= 0 && a[0] < len(grid) && a[1] >= 0 && a[1] < len(grid[0])
}
func (a Coordinate) Unpack() (int, int) {
	return a[0], a[1]
}
