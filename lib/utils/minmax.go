package utils

import "math"

// Max ...
func Max(a, b, c float32) float32 {
	at, bt, ct := float64(a), float64(b), float64(c)
	m := math.Max(at, bt)
	return float32(math.Max(m, ct))

}

// Min ...
func Min(a, b, c float32) float32 {
	at, bt, ct := float64(a), float64(b), float64(c)
	m := math.Min(at, bt)
	return float32(math.Min(m, ct))
}
