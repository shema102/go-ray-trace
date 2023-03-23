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

func SavePPMImageLineByLine(filename string, width, height int, c chan *[]renderer.Color) error {
	widthHeight := strconv.Itoa(width) + " " + strconv.Itoa(height)

	header := "P3\n" + widthHeight + "\n255\n"

	image := header

	fmt.Printf("Saving image to %s\n", filename)

	file, fileErr := os.Create(filename)

	if fileErr != nil {
		return fileErr
	}

	_, writeErr := file.WriteString(image)

	if writeErr != nil {
		return writeErr
	}

	rowNumber := 1

	for {
		row := <-c

		if row == nil {
			break
		}

		fmt.Printf("Saving row %d of %d\n", rowNumber, height)

		for j := 0; j < width; j++ {
			red := clip((*row)[j].X * 255.99)
			green := clip((*row)[j].Y * 255.99)
			blue := clip((*row)[j].Z * 255.99)

			_, writeErr := file.WriteString(strconv.Itoa(red) + " " + strconv.Itoa(green) + " " + strconv.Itoa(blue) + " ")

			if writeErr != nil {
				return writeErr
			}
		}

		rowNumber++
	}

	saveErr := file.Close()

	if saveErr != nil {
		return saveErr
	}

	fmt.Printf("Image saved to %s\n", filename)

	return nil
}
