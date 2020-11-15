package utils

import (
	"image"
)

// GetImageDim ...
func GetImageDim(img *image.Gray) (int, int) {
	b := img.Bounds()
	return b.Max.X, b.Max.Y
}

// GetDimensions ...
func GetDimensions(img *image.Gray) (int, int) {
	max := img.Bounds().Max
	return max.X, max.Y
}
