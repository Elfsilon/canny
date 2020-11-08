package resize

import (
	"github.com/nfnt/resize"
	"image"
)

// Scale returns scaled copy of given image
func Scale(width, height uint, img image.Image) image.Image {
	return resize.Resize(width, height, img, resize.MitchellNetravali)
}
