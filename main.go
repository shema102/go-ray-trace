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
	width := 820
	height := int(float64(width) / aspectRatio)
	samplesPerPixel := 20
	maxDepth := 8

	// World
	materialGround := t.Lambertian{Albedo: t.Color{X: 0.5, Y: 0.5, Z: 0.5}}

	world := t.CreateWorld()
	world.Add(t.Sphere{Center: t.Point3{X: 0, Y: -1000, Z: -1}, Radius: 1000, Material: materialGround})

	for a := -20; a < 20; a++ {
		for b := -20; b < 20; b++ {
			center := t.Point3{X: float64(a) + 0.9*util.Random(), Y: 0.2, Z: float64(b) + 0.9*util.Random()}
			selectMaterial := util.Random()

			randomSize := util.RandomRange(0.2, 0.3)

			if center.Sub(t.Point3{X: 4, Y: 0.2, Z: 0}).Length() > 0.9 {
				var material t.Material
				if selectMaterial < 0.5 {
					// diffuse
					albedo := t.RandomVec3().Mul(t.RandomVec3())
					material = t.Lambertian{Albedo: albedo}
				} else if selectMaterial < 0.8 {
					// metal
					albedo := t.RandomVec3().MulScalar(0.5).Add(t.Color{X: 0.5, Y: 0.5, Z: 0.5})
					fuzz := util.RandomRange(0, 0.5)
					material = t.Metal{Albedo: albedo, Fuzz: fuzz}
				} else {
					// glass
					material = t.Dielectric{RefractiveIndex: 1.5, Albedo: t.Color{X: 1, Y: 1, Z: 1}}
				}

				world.Add(t.Sphere{Center: center, Radius: randomSize, Material: material})
			}
		}
	}

	world.Add(t.Sphere{Center: t.Point3{X: 0, Y: 1, Z: 0}, Radius: 1.0, Material: t.Dielectric{RefractiveIndex: 1.5, Albedo: t.Color{X: 1, Y: 1, Z: 1}}})
	world.Add(t.Sphere{Center: t.Point3{X: -4, Y: 1, Z: 0}, Radius: 1.0, Material: t.Lambertian{Albedo: t.Color{X: 0.4, Y: 0.2, Z: 0.1}}})
	world.Add(t.Sphere{Center: t.Point3{X: 4, Y: 1, Z: 0}, Radius: 1.0, Material: t.Metal{Albedo: t.Color{X: 0.7, Y: 0.6, Z: 0.5}, Fuzz: 0.0}})

	// Camera
	camera := t.NewCamera(t.Point3{X: 13, Y: 2, Z: 3}, t.Point3{X: 0, Y: 0, Z: 0}, t.Vec3{X: 0, Y: 1, Z: 0}, 30, aspectRatio)

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
