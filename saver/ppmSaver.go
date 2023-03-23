package saver

import (
	"fmt"
	"os"
	"rt/renderer"
	"strconv"
)

func clip(x float64) int {
	if x < 0 {
		return 0
	} else if x > 255 {
		return 255
	} else {
		return int(x)
	}
}

func SavePPM(filename string, pixels [][]renderer.Vec3) error {
	width := len(pixels)
	height := len(pixels[0])

	widthHeight := strconv.Itoa(width) + " " + strconv.Itoa(height)

	header := "P3\n" + widthHeight + "\n255\n"

	image := header

	fmt.Printf("Saving image to %s\n", filename)

	for i := 0; i < width; i++ {
		fmt.Printf("Saving row %d of %d\n", i, width)
		for j := 0; j < height; j++ {
			r := clip(pixels[i][j].X * 255.99)
			g := clip(pixels[i][j].Y * 255.99)
			b := clip(pixels[i][j].Z * 255.99)

			image += strconv.Itoa(r) + " " + strconv.Itoa(g) + " " + strconv.Itoa(b) + "\n"
		}
	}

	file, fileErr := os.Create(filename)

	if fileErr != nil {
		return fileErr
	}

	_, writeErr := file.WriteString(image)

	if writeErr != nil {
		return writeErr
	}

	saveErr := file.Close()

	if saveErr != nil {
		return saveErr
	}

	fmt.Printf("Image saved to %s\n", filename)

	return nil
}
