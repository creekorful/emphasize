package main

import (
	"github.com/fogleman/gg"
	"log"
)

const (
	width    = 1920
	height   = 1080
	quote    = "At least there's stars in the sky."
	fontPath = "/Library/Fonts/Arial Unicode.ttf"
)

func main() {
	// Create random full HD image
	dc := gg.NewContext(width, height)

	// Create random black image
	dc.SetRGB255(0, 0, 0)
	dc.Clear()

	if err := drawQuote(dc, quote); err != nil {
		log.Fatal(err)
	}

	if err := dc.SavePNG("out.png"); err != nil {
		log.Fatal(err)
	}
}

func drawQuote(ctx *gg.Context, quote string) error {
	width := float64(ctx.Width())
	height := float64(ctx.Height())

	// TODO: compute font points based on image size and text length
	if err := ctx.LoadFontFace(fontPath, 96); err != nil {
		return err
	}

	ctx.SetRGB255(241, 196, 12)
	ctx.DrawStringAnchored(quote, width/2, height-(height/6), 0.5, 0.5)

	return nil
}
