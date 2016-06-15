package main

import (
	"image/color"
	"image"
	"os"
	"image/png"
)

type myImage struct {}

func (i myImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i myImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1024, 1024)
}

func (i myImage) At(x, y int) color.Color {
	blue := 0
	red := 0
	red = x

	myRGBA := color.RGBA{A: 255, B: uint8(blue), R: uint8(red)}
	return myRGBA
}


func main() {
	file, _ := os.Create("/tmp/myImage.png")

	curImage := myImage{}

	png.Encode(file, curImage)
}
