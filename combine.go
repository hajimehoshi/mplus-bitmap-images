package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

func main() {
	const charFullWidth = 12
	const charHalfWidth = 6
	const charHeight = 16

	files := []string{
		"latin.png",
		"bmp-0.png",
		"bmp-2.png",
		"bmp-3.png",
		"bmp-4.png",
		"bmp-5.png",
		"bmp-6.png",
		"bmp-7.png",
		"bmp-8.png",
		"bmp-9.png",
		"bmp-15.png",
	}
	palette := color.Palette([]color.Color{
		color.Transparent, color.Opaque,
	})
	result := image.NewPaletted(image.Rect(0, 0, 256*charFullWidth, 256*charHeight), palette)
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		if file == "latin.png" {
			for i := 0; i < 256; i++ {
				dx := i * charFullWidth
				dy := 0
				dr := image.Rect(dx, dy, dx + charHalfWidth, dy + charHeight)
				sp := image.Point{
					(i % 32) * charHalfWidth,
					(i / 32) * charHeight,
				}
				draw.Draw(result, dr, img, sp, draw.Src)
			}
			continue
		}
		id, err := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(file))
		if err != nil {
			panic(err)
		}
		for i := 0; i < 4096; i++ {
			if id == 0 && i < 256 {
				continue
			}
			dx := (i % 256) * charFullWidth
			dy := id * 256 + (i / 256) * charHeight
			dr := image.Rect(dx, dy, dx + charFullWidth, dy + charHeight)
			sp := image.Point{
				(i % 64) * charFullWidth,
				(i / 64) * charHeight,
			}
			draw.Draw(result, dr, img, sp, draw.Src)
		}
	}
	e := png.Encoder{CompressionLevel: png.BestCompression}
	if err := e.Encode(os.Stdout, result); err != nil {
		panic(err)
	}
}
