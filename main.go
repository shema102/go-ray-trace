package main

import (
	"fmt"
	ren "rt/renderer"
	"rt/saver"
	"sync"
)

var wg sync.WaitGroup

func main() {
	filename := "test.ppm"

	// Image
	aspectRatio := 16.0 / 9.0
	width := 400
	height := int(float64(width) / aspectRatio)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := ren.Vec3{X: 0, Y: 0, Z: 0}
	horizontal := ren.Vec3{X: viewportWidth, Y: 0, Z: 0}
	vertical := ren.Vec3{X: 0, Y: viewportHeight, Z: 0}
	lowerLeftCorner := origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(ren.Vec3{X: 0, Y: 0, Z: focalLength})

	c := make(chan *[]ren.Color)
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := saver.SavePPMImageLineByLine(filename, width, height, c)
		if err != nil {
			fmt.Printf("Error during image save: %s", err)
			return
		}
	}()

	fmt.Printf("Rendering image to %s\n", filename)

	for j := height - 1; j >= 0; j-- {
		line := make([]ren.Color, width)

		fmt.Printf("Saving row %d of %d\n", height-j, height)

		for i := 0; i < width; i++ {
			u := float64(i) / float64(width-1)
			v := float64(j) / float64(height-1)

			ray := ren.Ray{Origin: origin, Direction: lowerLeftCorner.Add(horizontal.MulScalar(u)).Add(vertical.MulScalar(v)).Sub(origin)}
			line[i] = ray.RayColor()
		}

		c <- &line
	}

	close(c)

	wg.Wait()

	fmt.Printf("Image saved to %s\n", filename)
}
