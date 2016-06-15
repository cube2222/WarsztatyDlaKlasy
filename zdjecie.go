package main

import (
	"image/color"
	"image"
	"os"
	"fmt"
	"image/png"
)

type myImage struct {
}

func (img myImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img myImage) Bounds() image.Rectangle {
	return image.Rect(0,0,1023,1023)
}

func (img myImage) At(x int, y int) color.Color {
	blue := uint8((x+y)*(x+y)-x*y)
	red := uint8((x+20+y)*(x+20+y)-(20+x)*y)
	return color.RGBA{A: 255, R: red, G: 0, B: blue}
}

func main() {
	img := myImage{}

	file, err := os.Create("C:/tmp/myImage.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	png.Encode(file, img)
}
