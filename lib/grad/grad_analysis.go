package grad

import (
	"github.com/Elfsilon/canny/lib/utils"
	"image"
)

func gradAnalysis(img *image.Gray, op *Operator) ([][]float64, [][]uint8) {
	_, width, height := utils.CreateImage(img)
	gradMat := make([][]float64, height)
	angMat := make([][]uint8, height)
	for i := range gradMat {
		gradMat[i] = make([]float64, width)
		angMat[i] = make([]uint8, width)
	}

	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			G, ang := op.grad(x, y, img)
			gradMat[y][x] = G
			angMat[y][x] = ang
		}
	}

	return gradMat, angMat
}
