package main

import (
	"image/color"
	"image"
	"fmt"
	"os"
	"image/png"
	"math/cmplx"
)

var i int

type Mandelbrot struct {
	colorMax color.RGBA
	iterationCountOnPixels [][]int
	iterationCount int
	width int
	height int
}

func (m Mandelbrot) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Mandelbrot) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

func (m *Mandelbrot) At(x int, y int) color.Color {
	iteration := float64(m.iterationCountOnPixels[x][y])
	iterMax := float64(m.iterationCount)
	red := uint8(((float64(m.colorMax.R)/iterMax)*iteration))
	green := uint8(((float64(m.colorMax.G)/iterMax)*iteration))
	blue := uint8(((float64(m.colorMax.B)/iterMax)*iteration))
	return color.RGBA{red, green, blue, 255}
}

func (m *Mandelbrot) calculateMandelbrot() {

	width := m.width
	height := m.height
	for row := 0; row < width; row++ {
		for col := 0; col < height; col++ {
			c := complex((float64(col-width) / 2.0) * 4.0 / float64(width), (float64(row-height) / 2.0) * 4.0 / float64(width))
			z := complex(0, 0)
			iteration := 0
			for cmplx.Abs(z) <= 2*2 && iteration < m.iterationCount {
				x := real(z)*real(z) - imag(z)*imag(z) + real(c)
				y := 2 * real(z) * imag(z) + imag(c)
				z = complex(x, y)
				iteration++
				m.iterationCountOnPixels[col][row] += 1
			}
		}
	}
}

func New(colorMax color.RGBA, iterationCount int, width int, height int) *Mandelbrot {
	iterationCountOnPixels := make([][]int, width, width)
	for index, _ := range iterationCountOnPixels {
		column := make([]int, height, height)
		iterationCountOnPixels[index] = column
	}
	return &Mandelbrot{colorMax, iterationCountOnPixels, iterationCount, width, height}
}

// add color min and color max
func main() {
	myMandelbrot := New(color.RGBA{255, 0, 0, 255}, 1000, 2048, 2048)
	myMandelbrot.calculateMandelbrot()
	file, _ := os.Create("C:/tmp/mandelbrotImage.png")
	fmt.Println("Encoding")
	png.Encode(file, myMandelbrot)
	fmt.Println("Finished encoding")
}
