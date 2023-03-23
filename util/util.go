package util

import (
	"math"
	"math/rand"
)

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	} else if x > max {
		return max
	} else {
		return x
	}
}

func Random() float64 {
	return rand.Float64()
}

func RandomRange(min, max float64) float64 {
	return min + (max-min)*Random()
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}
