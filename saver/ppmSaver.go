package saver

import (
	"math"
	"os"
	t "rt/tracer"
	"rt/util"
	"strconv"
)

func SavePPMImageLineByLine(filename string, width, height int, scale float64, c chan *[]t.Color) error {
	widthHeight := strconv.Itoa(width) + " " + strconv.Itoa(height)

	header := "P3\n" + widthHeight + "\n255\n"

	image := header

	file, fileErr := os.Create(filename)

	if fileErr != nil {
		return fileErr
	}

	_, writeErr := file.WriteString(image)

	if writeErr != nil {
		return writeErr
	}

	for {
		row := <-c

		if row == nil {
			break
		}

		for j := 0; j < width; j++ {
			red := (*row)[j].X
			green := (*row)[j].Y
			blue := (*row)[j].Z

			red = math.Sqrt(red * scale)
			green = math.Sqrt(green * scale)
			blue = math.Sqrt(blue * scale)

			rowStr := strconv.Itoa(int(256*util.Clamp(red, 0, 0.999))) + " " +
				strconv.Itoa(int(256*util.Clamp(green, 0, 0.999))) + " " +
				strconv.Itoa(int(256*util.Clamp(blue, 0, 0.999))) + "\n"

			_, writeErr := file.WriteString(rowStr)

			if writeErr != nil {
				return writeErr
			}
		}
	}

	saveErr := file.Close()

	if saveErr != nil {
		return saveErr
	}

	return nil
}
