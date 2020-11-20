package recognition

import (
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/Elfsilon/canny/lib/utils"
)

// NewBaseImage ...
func NewBaseImage(path string, ratio []float64) BaseImage {
	bimage := BaseImage{
		path:  path,
		ratio: ratio,
	}
	bimage.loadAndGray()
	return bimage
}

// BaseImage ...
// []Ratio - coefficients of digits (index for digit - 8 for 8 digit)
type BaseImage struct {
	path  string
	ratio []float64
	img   image.Gray
}

// Determine ...
func (r *BaseImage) loadAndGray() {
	imgFile, err := os.Open(r.path)
	if err != nil {
		log.Println("error", err)
	}
	defer imgFile.Close()

	loadedImg, _ := jpeg.Decode(imgFile)
	gray := utils.ToGrayscale(loadedImg)

	r.img = *gray
}

// NewRecognizer ...
func NewRecognizer(base []BaseImage) Recognizer {
	return Recognizer{
		base: base,
	}
}

// Recognizer ...
type Recognizer struct {
	base []BaseImage
}

type variant struct {
	index int
	value float64
}

// Determine ...
func (r *Recognizer) Determine(img *image.Gray, method string) (int, *image.Gray) {
	var variants []variant
	for _, baseimg := range r.base {
		var k float64
		var err error

		switch method {
		case "euclidean":
			k, err = utils.CompareEuclid(&baseimg.img, img)
		case "manhattan":
			k, err = utils.CompareManhattan(&baseimg.img, img)
		case "ncc":
			k, err = utils.CompareNCC(&baseimg.img, img)
		default:
			k, err = utils.Compare(&baseimg.img, img)
		}

		if err != nil {
			log.Println(err)
		}

		var max float64 = -1
		var index int = 0 // Index keeps possible result digit which we recognizing

		for i, r := range baseimg.ratio {
			if m := r * k; m > max {
				max = m
				index = i
			}
		}

		variants = append(variants, variant{
			index: index,
			value: max,
		})
	}

	var max float64 = -1
	var index int = 0 // Index keeps end result digit which we recognizing

	for _, v := range variants {
		if v.value > max {
			max = v.value
			index = v.index
		}
	}

	return index, &r.base[index].img
}
