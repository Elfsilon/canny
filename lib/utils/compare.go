package utils

import (
	"errors"
	"image"
)

// Compare ...
func Compare(img1 *image.Gray, img2 *image.Gray) (float64, error) {
	width, height := GetDimensions(img1)
	width2, height2 := GetDimensions(img2)

	if width != width2 || height != height2 {
		return -1, errors.New("Cannot compare pictures with different dimensions")
	}

	summary := width * height
	sameCount := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if GetGray(x, y, img1) == GetGray(x, y, img2) {
				sameCount++
			}
		}
	}

	return float64(sameCount) / float64(summary), nil
}
