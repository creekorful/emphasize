package main

import "github.com/fogleman/gg"

const (
	width  = 1920
	height = 1080
)

func main() {
	dc := gg.NewContext(width, height)

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/Library/Fonts/Arial Unicode.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", width/2, height/2, 0.5, 0.5)
	_ = dc.SavePNG("out.png")
}
