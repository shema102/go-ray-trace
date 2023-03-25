package main

import (
	"fmt"
	"rt/saver"
	t "rt/tracer"
	"rt/util"
	"sync"
)

var wg sync.WaitGroup

func main() {
	filename := "test.ppm"

	// Image
	aspectRatio := 16.0 / 9.0
	width := 1000
	height := int(float64(width) / aspectRatio)
	samplesPerPixel := 10
	maxDepth := 5

	// World
	materialGround := t.Lambertian{Albedo: t.Color{X: 0.8, Y: 0.8, Z: 0.0}}
	materialCenter := t.Lambertian{Albedo: t.Color{X: 0.1, Y: 0.2, Z: 0.5}}
	materialLeft := t.Dielectric{RefractiveIndex: 1.5, Albedo: t.RGBToColor("#d238d1")}
	materialRight := t.Metal{Albedo: t.Color{X: 0.8, Y: 0.6, Z: 0.2}, Fuzz: 0.5}

	world := t.CreateWorld()
	world.Add(t.Sphere{Center: t.Point3{X: 0, Y: -100.5, Z: -1}, Radius: 100, Material: materialGround})
	world.Add(t.Sphere{Center: t.Point3{X: 0, Y: 0, Z: -1}, Radius: 0.5, Material: materialCenter})
	world.Add(t.Sphere{Center: t.Point3{X: -1, Y: 0, Z: -1}, Radius: 0.5, Material: materialLeft})
	world.Add(t.Sphere{Center: t.Point3{X: 1, Y: 0, Z: -1}, Radius: 0.5, Material: materialRight})

	// Camera
	camera := t.NewCamera(120, 16/9)

	c := make(chan *[]t.Color)
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := saver.SavePPMImageLineByLine(filename, width, height, 1.0/float64(samplesPerPixel), c)
		if err != nil {
			fmt.Printf("Error during image save: %s", err)
			return
		}
	}()

	fmt.Printf("Rendering image to %s\n", filename)

	startTime := util.Now()

	for j := height - 1; j >= 0; j-- {
		line := make([]t.Color, width)

		fmt.Printf("Rendering row %d of %d\n", height-j, height)

		for i := 0; i < width; i++ {
			pixelColor := t.Color{X: 0, Y: 0, Z: 0}

			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + util.Random()) / float64(width-1)
				v := (float64(j) + util.Random()) / float64(height-1)

				ray := camera.GetRay(u, v)
				pixelColor = pixelColor.Add(t.TraceRay(ray, world, maxDepth))
			}

			line[i] = pixelColor
		}

		c <- &line
	}

	close(c)

	wg.Wait()

	fmt.Printf("Image saved to %s\n", filename)
	fmt.Printf("Render time: %f\n", util.Since(startTime))
}
