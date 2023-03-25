package util

import (
	"math"
	"math/rand"
	"strconv"
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

func HexToRGB(hex string) (int, int, int) {
	r, _ := strconv.ParseInt(hex[1:3], 16, 64)
	g, _ := strconv.ParseInt(hex[3:5], 16, 64)
	b, _ := strconv.ParseInt(hex[5:7], 16, 64)
	return int(r), int(g), int(b)
}
