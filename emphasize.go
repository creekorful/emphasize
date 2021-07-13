package main

import "github.com/fogleman/gg"

const (
	width    = 1920
	height   = 1080
	sentence = "At least there's stars in the sky."
)

func main() {
	dc := gg.NewContext(width, height)

	dc.SetRGB255(0, 0, 0)
	dc.Clear()

	dc.SetRGB255(241, 196, 12)
	if err := dc.LoadFontFace("/Library/Fonts/Arial Unicode.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(sentence, width/2, height/2, 0.5, 0.5)
	_ = dc.SavePNG("out.png")
}
