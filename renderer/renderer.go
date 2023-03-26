package renderer

import (
	"fmt"
	t "rt/tracer"
	"rt/util"
)

type RenderResult struct {
	rowNumber int
	line      []t.Color
}

func RenderChunkWorker(in <-chan int, out chan<- RenderResult, camera *t.Camera, world *t.World, samplesPerPixel, width, height, depth int) {
	for row := range in {
		line := make([]t.Color, width)

		fmt.Printf("Rendering line %d\n", row)

		for i := 0; i < width; i++ {
			col := t.Color{}

			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + util.RandomFloat()) / float64(width)
				v := (float64(row) + util.RandomFloat()) / float64(height)

				r := camera.GetRay(u, v)
				col = col.Add(t.TraceRay(r, world, depth))
			}

			// apply gamma correction
			col = col.DivScalar(float64(samplesPerPixel))
			col = col.Sqrt()
			line[i] = col
		}

		out <- RenderResult{
			rowNumber: row,
			line:      line,
		}

		fmt.Printf("Done rendering line %d\n", row)
	}
}

func RenderImage(width int, aspectRatio float64, camera *t.Camera, world *t.World, samplesPerPixel, depth int) [][]t.Color {
	const maxWorkers = 40

	height := int(float64(width) / aspectRatio)

	in := make(chan int, height)
	out := make(chan RenderResult, height)

	for i := 0; i < maxWorkers; i++ {
		go RenderChunkWorker(in, out, camera, world, samplesPerPixel, width, height, depth)
	}

	chunks := make([][]t.Color, height)

	for j := 0; j < height; j++ {
		in <- j
	}

	for i := 0; i < height; i++ {
		result := <-out
		chunks[height-result.rowNumber-1] = result.line
	}

	return chunks
}
