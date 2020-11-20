package utils

import (
	"errors"
	"image"
	"math"
)

func checkDimensions(w1, h1, w2, h2 int) error {
	if w1 != w2 || h1 != h2 {
		return errors.New("Cannot compare pictures with different dimensions")
	}
	return nil
}

// Compare ...
func Compare(img1 *image.Gray, img2 *image.Gray) (float64, error) {
	width, height := GetDimensions(img1)
	width2, height2 := GetDimensions(img2)

	err := checkDimensions(width, height, width2, height2)
	if err != nil {
		return -1, err
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

// CompareEuclid - compares 2 images using euclidean norm and return ratio
func CompareEuclid(img1 *image.Gray, img2 *image.Gray) (float64, error) {
	width, height := GetDimensions(img1)
	width2, height2 := GetDimensions(img2)

	err := checkDimensions(width, height, width2, height2)
	if err != nil {
		return -1, err
	}

	summary := width * height
	var sum float64 = 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sum += math.Pow(GetGray(x, y, img1)-GetGray(x, y, img2), 2)
		}
	}

	return math.Sqrt(sum) / float64(summary), nil
}

// CompareManhattan - compares 2 images using manhattan norm and return ratio
func CompareManhattan(img1 *image.Gray, img2 *image.Gray) (float64, error) {
	width, height := GetDimensions(img1)
	width2, height2 := GetDimensions(img2)

	err := checkDimensions(width, height, width2, height2)
	if err != nil {
		return -1, err
	}

	summary := width * height
	var sum float64 = 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sum += math.Abs(GetGray(x, y, img1) - GetGray(x, y, img2))
		}
	}

	return sum / float64(summary), nil
}

// CompareNCC ...
func CompareNCC(img1 *image.Gray, img2 *image.Gray) (float64, error) {
	width, height := GetDimensions(img1)
	width2, height2 := GetDimensions(img2)

	err := checkDimensions(width, height, width2, height2)
	if err != nil {
		return -1, err
	}

	summary := width * height
	var mean1, mean2 float64 = 0, 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			mean1 += GetGray(x, y, img1)
			mean2 += GetGray(x, y, img2)
		}
	}
	mean1, mean2 = mean1/float64(summary), mean2/float64(summary)

	var sum, std1, std2 float64 = 0, 0, 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sum += (GetGray(x, y, img1) - mean1) * (GetGray(x, y, img2) - mean2)
			std1 += math.Pow(GetGray(x, y, img1)-mean1, 2)
			std2 += math.Pow(GetGray(x, y, img2)-mean2, 2)
		}
	}
	std1 = math.Sqrt(std1 / (float64(summary) - 1))
	std2 = math.Sqrt(std2 / (float64(summary) - 1))

	return sum / (float64(summary) * std1 * std2), nil
}
