package main

import (
	"image/color"
	"image"
	"image/png"
	"net/http"
	"net/url"
	"strconv"
)

type Pixel struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

type MyImage struct {
	tab [1024][1024]Pixel
}

func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1024, 1024)
}

func (img MyImage) At(x int, y int) color.Color {
	return color.RGBA{A: img.tab[x][y].a, R: img.tab[x][y].r, B: img.tab[x][y].b, G: img.tab[x][y].g}
}

func (img MyImage) calculateMandelbrot() MyImage {
	width := 1023
	height := 1023
	for row := 0; row < 1023; row++ {
		for col := 0; col < 1023; col++ {
			c_re := float64((col - width / 2.0) * 4.0 / width)
			c_im := float64((row - height / 2.0) * 4.0 / width)
			x := float64(0)
			y := float64(0)
			iteration := 0
			for x * x + y * y <= 4 && iteration < 100 {
				x_new := x * x - y * y + c_re
				y = 2 * x * y + c_im
				x = x_new
				iteration++
			}
			if (iteration < 100) {
				img.tab[col][row] = Pixel{255, 255, 255, 255}
			} else {
				img.tab[col][row] = Pixel{0, 0, 0, 0}
			}
		}
	}
	return img
}

func main() {
	img := MyImage{}
	img = img.calculateMandelbrot()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	})
	http.ListenAndServe(":80", nil)
}
