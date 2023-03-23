package main

import (
	"rt/renderer"
	"rt/saver"
)

func main() {

	image := make([][]renderer.Vec3, 512)

	for i := 0; i < 512; i++ {
		image[i] = make([]renderer.Vec3, 512)
		for j := 0; j < 512; j++ {
			image[i][j] = renderer.Vec3{X: float64(i) / 512.0, Y: float64(j) / 512.0, Z: 0.25}
		}
	}

	err := saver.SavePPM("test.ppm", image)
	if err != nil {
		return
	}
}
