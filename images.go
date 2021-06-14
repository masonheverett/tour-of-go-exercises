package main

import (
	"image"
	"image/color"

	pic "golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x + y) % 256), uint8((x + (2 * y)) % 256), uint8((x + (3 * y)) % 256), 255}
}

func Images() {
	m := Image{}
	pic.ShowImage(m)
}
