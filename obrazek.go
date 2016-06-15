package main

import (
	"image/color"
	"image"
	"image/png"
	"net/http"
	"net/url"
	"strconv"
)

type MyImage struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1024, 1024)
}

func (img MyImage) At(x int, y int) color.Color {
	return color.RGBA{A: img.a, R: img.r, B: img.b, G: img.g}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params, _ := url.ParseQuery(r.URL.RawQuery)
		alpha, _ := strconv.Atoi(params["a"][0])
		red, _ := strconv.Atoi(params["r"][0])
		green, _ := strconv.Atoi(params["g"][0])
		blue, _ := strconv.Atoi(params["b"][0])
		img := MyImage{
			a: uint8(alpha),
			r: uint8(red),
			g: uint8(green),
			b: uint8(blue),
		}
		png.Encode(w, img)
	})
	http.ListenAndServe(":80", nil)
}
