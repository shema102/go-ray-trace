package util

import "time"

func Now() float64 {
	return float64(time.Now().UnixNano()) / 1e9
}

func Since(start float64) float64 {
	return Now() - start
}
