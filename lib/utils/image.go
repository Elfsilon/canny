package utils

import (
	"image"
)

// GetGray ...
func GetGray(x, y int, img *image.Gray) float64 {
	col := img.GrayAt(x, y)
	return float64(col.Y)
}

// CreateImage ...
func CreateImage(img *image.Gray) (*image.Gray, int, int) {
	max := img.Bounds().Max
	min := img.Bounds().Min
	return image.NewGray(image.Rect(max.X, max.Y, min.X, min.Y)), max.X, max.Y
}

// ToGrayscale ...
func ToGrayscale(img image.Image) *image.Gray {
	max := img.Bounds().Max
	min := img.Bounds().Min
	gray := image.NewGray(image.Rect(max.X, max.Y, min.X, min.Y))

	for y := 0; y < max.Y; y++ {
		for x := 0; x < max.X; x++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	return gray
}
