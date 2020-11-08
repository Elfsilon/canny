package grayscale

import (
	"image"
	"image/color"
	"github.com/Elfsilon/canny/lib/utils"
)

// ComputeAVG ...
func ComputeAVG(r, g, b float32) uint8 {
	return uint8((r + g + b) / 3)
}

// ComputeLuminosity ...
func ComputeLuminosity(r, g, b float32) uint8 {
	return uint8(0.07*b + 0.72*g + 0.21*r)
}

// ComputeLightness ...
func ComputeLightness(r, g, b float32) uint8 {
	return uint8((utils.Max(r, g, b) + utils.Min(r, g, b)) / 2)
}

func createGray(width, height int) *image.Gray {
	upLeft := image.Point{0, 0}
	bottomRight := image.Point{width, height}
	return image.NewGray(image.Rectangle{upLeft, bottomRight})
}

// Gray ...
func Gray(img image.Image, computeColor func(r, g, b float32) uint8) *image.Gray {
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	grayscaled := createGray(width, height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			nr, ng, nb := float32(r)/255, float32(g)/255, float32(b)/255
			graycolor := computeColor(nr, ng, nb)
			grayscaled.SetGray(x, y, color.Gray{uint8(graycolor)})
		}
	}

	return grayscaled
}
