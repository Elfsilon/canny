package grad

func findGradNeighbours(x, y int, ang uint8) (int, int, int, int) {
	if ang == 45 {
		return x + 1, y - 1, x - 1, y + 1
	} else if ang == 90 {
		return x, y + 1, x, y - 1
	} else if ang == 135 {
		return x - 1, y - 1, x + 1, y + 1
	} else {
		return x + 1, y, x - 1, y
	}
}

func nonmaxSuppression(gradMat [][]float64, dirMat [][]uint8) [][]float64 {
	height, width := len(gradMat), len(gradMat[0])
	resMat := make([][]float64, height)
	for i := range resMat {
		resMat[i] = make([]float64, width)
	}
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			xn1, yn1, xn2, yn2 := findGradNeighbours(x, y, dirMat[y][x])
			if gradMat[y][x] >= gradMat[yn1][xn1] && gradMat[y][x] >= gradMat[yn2][xn2] {
				resMat[y][x] = gradMat[y][x]
			} else {
				resMat[y][x] = 0
			}
		}
	}
	return resMat
}
