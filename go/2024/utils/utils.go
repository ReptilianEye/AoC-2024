package utils

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			print(cell)
		}
		println()
	}
}
