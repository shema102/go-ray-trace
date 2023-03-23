package main

import (
	"fmt"
	"rt/renderer"
	"rt/saver"
	"sync"
)

var wg sync.WaitGroup

func main() {
	height, width := 1024, 1024

	c := make(chan *[]renderer.Color)

	wg.Add(1)

	go func() {
		defer wg.Done()

		err := saver.SavePPMImageLineByLine("test.ppm", width, height, c)
		if err != nil {
			fmt.Printf("Error during image save: %s", err)
			return
		}
	}()

	for i := 0; i < height; i++ {
		line := make([]renderer.Color, width)
		for j := 0; j < width; j++ {
			line[j] = renderer.Color{X: float64(i) / float64(height), Y: float64(j) / float64(width), Z: 1}
		}
		c <- &line
	}

	close(c)

	wg.Wait()
}
