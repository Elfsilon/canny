package grad

func doubleThresold(mat [][]float64, lowerP, upperP float64) [][]float64 {
	height, width := len(mat), len(mat[0])
	lowerThresold := lowerP * 255
	upperThresold := upperP * 255
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if mat[y][x] >= upperThresold {
				mat[y][x] = 255
			} else if mat[y][x] <= lowerThresold {
				mat[y][x] = 0
			} else {
				mat[y][x] = 127
			}
		}
	}
	return mat
}
