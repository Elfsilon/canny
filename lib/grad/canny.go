package grad

import (
	// "fmt"
	"github.com/Elfsilon/canny/lib/utils"
	"image"
	"image/color"
)

// Canny operator implementation
func Canny(img *image.Gray, op *Operator, upThresold, downThresold float64) *image.Gray {
	gradMat, angMat := gradAnalysis(img, op)
	suppresed := nonmaxSuppression(gradMat, angMat)
	thresolded := doubleThresold(suppresed, downThresold, upThresold)
	tracked := blobAnalysis(thresolded, 255, 127, 0)

	resImage, width, height := utils.CreateImage(img)
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			resImage.SetGray(x, y, color.Gray{Y: uint8(tracked[y][x])})
		}
	}

	return resImage
}
