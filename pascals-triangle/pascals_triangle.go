package pascal

func Triangle(rows int) [][]int {
	output := make([][]int, rows)

	for row := 0; row < rows; row++ {
		output[row] = make([]int, row+1)
		for col := 0; col <= row; col++ {
			if col == 0 || col == row {
				output[row][col] = 1
				continue
			}
			output[row][col] = output[row-1][col-1] + output[row-1][col]
		}
	}

	return output
}
