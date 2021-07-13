package main

import (
	"github.com/fogleman/gg"
	"log"
	"math/rand"
	"os"
)

const (
	quote    = "At least there's stars in the sky."
	fontPath = "/Library/Fonts/Arial Unicode.ttf"
)

func main() {
	im, err := gg.LoadImage(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContextForImage(im)

	if err := applyNoise(dc, 32.0+128.0); err != nil {
		log.Fatal(err)
	}

	if err := drawQuote(dc, quote); err != nil {
		log.Fatal(err)
	}

	if err := dc.SavePNG("out.png"); err != nil {
		log.Fatal(err)
	}
}

func applyNoise(ctx *gg.Context, sigma float64) error {
	width, height := ctx.Width(), ctx.Height()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Gaussian noise
			noise := uint32(rand.NormFloat64() * sigma)

			r, g, b, _ := ctx.Image().At(x, y).RGBA()

			ctx.SetRGBA255(int(r), int(g), int(b), int(noise))
			ctx.SetPixel(x, y)
		}
	}

	return nil
}

func drawQuote(ctx *gg.Context, quote string) error {
	width, height := float64(ctx.Width()), float64(ctx.Height())

	// TODO: compute font points based on image size and text length (default 96)
	if err := ctx.LoadFontFace(fontPath, 56); err != nil {
		return err
	}

	ctx.SetRGB255(241, 196, 12)

	ctx.DrawStringAnchored(quote, width/2, height-(height/10), 0.5, 0.5)

	return nil
}
