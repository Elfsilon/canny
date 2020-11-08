package grad

func blobAnalysis(mat [][]float64, max, high, clear float64) [][]float64 {
	height, width := len(mat), len(mat[0])
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if mat[y][x] == max {
				continue
			}
			if mat[y][x] == high && !isTouch(mat, x, y, max) {
				mat[y][x] = clear
			}
		}
	}
	return mat
}

func isTouch(mat [][]float64, x, y int, max float64) bool {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i != 0 && j != 0 && mat[y-i][x-j] == max {
				return true
			}
		}
	}
	return false
}
