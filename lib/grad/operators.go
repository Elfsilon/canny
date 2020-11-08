package grad

import (
	"github.com/Elfsilon/canny/lib/utils"
	"image"
	"math"
)

// Operator ...
type Operator struct {
	MGx [3][3]float64
	MGy [3][3]float64
}

// Roberts operator
func Roberts() *Operator {
	return &Operator{
		MGx: [3][3]float64{
			{1, 0, 0},
			{0, -1, 0},
			{0, 0, 0},
		},
		MGy: [3][3]float64{
			{0, 1, 0},
			{-1, 0, 0},
			{0, 0, 0},
		},
	}
}

// Sobel operator
func Sobel() *Operator {
	return &Operator{
		MGx: [3][3]float64{
			{-1, 0, 1},
			{-2, 0, 2},
			{-1, 0, 1},
		},
		MGy: [3][3]float64{
			{-1, -2, -1},
			{0, 0, 0},
			{1, 2, 1},
		},
	}
}

// Prewitt operator
func Prewitt() *Operator {
	return &Operator{
		MGx: [3][3]float64{
			{1, 0, -1},
			{1, 0, -1},
			{1, 0, -1},
		},
		MGy: [3][3]float64{
			{1, 1, 1},
			{0, 0, 0},
			{-1, -1, -1},
		},
	}
}

func roundAngle(ang float64) uint8 {
	if ang >= 22 && ang < 67 {
		return 45
	} else if ang < 112 {
		return 90
	} else if ang < 158 {
		return 135
	} else {
		return 0
	}
}

func (o *Operator) grad(x, y int, img *image.Gray) (float64, uint8) {
	var Gx, Gy float64 = 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val := utils.GetGray(x-1+i, y-1+j, img)
			Gx += val * o.MGx[i][j]
			Gy += val * o.MGy[i][j]
		}
	}
	G := math.Sqrt(Gx*Gx + Gy*Gy)
	angle := roundAngle(math.Abs(math.Atan2(Gy, Gx) * 180 / math.Pi))
	return G, angle
}
