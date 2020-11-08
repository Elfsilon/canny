package blur

import (
	"image"
	"math"
)

func gaussian2d(x, y, mu, sigma float64) float64 {
	return math.Exp(-(math.Pow(x-mu, 2)/sigma+math.Pow(y-mu, 2)/sigma)/2.0) / (2 * math.Pi * sigma * sigma)
}

func gaussian1d(x, sigma, mu float64) float64 {
	return math.Exp(-math.Pow((x-mu)/sigma, 2.0)/2.0) / (math.Sqrt(2*math.Pi) * sigma)
}

func setupKernel(r int, sigma float64) [][]float64 {
	var sum float64 = 0
	var kerW int = 2*r + 1
	kernel2d := make([][]float64, kerW)
	for i := range kernel2d {
		kernel2d[i] = make([]float64, kerW)
	}

	// Building weight matrix of kernel 2D
	for x := 0; x < kerW; x++ {
		for y := 0; y < kerW; y++ {
			// kernel2d[x][y] = kernel1d[x] * kernel1d[y]
			kernel2d[x][y] = gaussian2d(float64(x), float64(y), float64(r), sigma)
			sum += kernel2d[x][y]
		}
	}

	// Normalize
	for x := 0; x < kerW; x++ {
		for y := 0; y < kerW; y++ {
			kernel2d[x][y] /= sum
		}
	}
	return kernel2d
}

func convolution(x, y, radius, width, height int, img *image.Gray, kernel [][]float64) float64 {
	var sum float64 = 0
	for xi := x - radius; xi < x+radius+1; xi++ {
		for yi := y - radius; yi < y+radius+1; yi++ {
			if xi >= 0 && xi < height && yi >= 0 && yi < width {
				offset := img.PixOffset(yi, xi)
				sum += float64(img.Pix[offset]) * kernel[xi-(x-radius)][yi-(y-radius)]
			}
		}
	}
	return sum
}

// GaussianBlur ...
func GaussianBlur(img *image.Gray, radius int, sigma float64) {
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			kernel := setupKernel(radius, sigma)
			center := convolution(i, j, radius, width, height, img, kernel)
			offset := img.PixOffset(j, i)
			img.Pix[offset] = uint8(center)
		}
	}
}
