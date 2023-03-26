package main

import (
	"fmt"
	"rt/renderer"
	"rt/saver"
	t "rt/tracer"
	"rt/util"
)

func main() {
	filename := "test.ppm"

	// Image
	aspectRatio := 16.0 / 9.0
	width := 1920
	height := int(float64(width) / aspectRatio)
	samplesPerPixel := 200
	maxDepth := 80

	// World
	materialGround := t.Lambertian{Albedo: t.Color{X: 0.5, Y: 0.5, Z: 0.5}}

	world := t.CreateWorld()
	world.Add(t.Sphere{Center: t.Point3{X: 0, Y: -1000, Z: -1}, Radius: 1000, Material: materialGround})

	for a := -20; a < 20; a++ {
		for b := -20; b < 20; b++ {
			center := t.Point3{X: float64(a) + 0.9*util.RandomFloat(), Y: 0.2, Z: float64(b) + 0.9*util.RandomFloat()}
			selectMaterial := util.RandomFloat()

			randomSize := util.RandomFloatInRange(0.2, 0.3)

			if center.Sub(t.Point3{X: 4, Y: 0.2, Z: 0}).Length() > 0.9 {
				var material t.Material
				if selectMaterial < 0.5 {
					// diffuse
					albedo := t.RandomVec3().Mul(t.RandomVec3())
					material = t.Lambertian{Albedo: albedo}
				} else if selectMaterial < 0.8 {
					// metal
					albedo := t.RandomVec3().MulScalar(0.5).Add(t.Color{X: 0.5, Y: 0.5, Z: 0.5})
					fuzz := util.RandomFloatInRange(0, 0.5)
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

	startTime := util.Now()

	renderedImage := renderer.RenderImage(width, aspectRatio, &camera, &world, samplesPerPixel, maxDepth)

	err := saver.SavePPMImage(filename, width, height, &renderedImage)

	if err != nil {
		fmt.Printf("Error during image save: %s", err)
		return
	}

	fmt.Printf("Image saved to %s\n", filename)
	fmt.Printf("Render time: %f\n", util.Since(startTime))
}
