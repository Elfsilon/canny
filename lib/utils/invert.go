package utils

import (
	"fmt"
	"image"
	"image/color"
)

// Invert ...
func Invert(img *image.Gray) *image.Gray {
	newimg, width, height := CreateImage(img)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == 0 && x == 0 {
				fmt.Println(GetGray(x, y, img), 255-GetGray(x, y, img))
			}
			col := 255 - GetGray(x, y, img)
			newimg.SetGray(x, y, color.Gray{Y: uint8(col)})
		}
	}

	return newimg
}
