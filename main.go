package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func getRGBA(file string) (*image.RGBA, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	// https://stackoverflow.com/questions/31463756/convert-image-image-to-image-nrgba
	// https://stackoverflow.com/questions/47535474/convert-image-from-image-ycbcr-to-image-rgba
	if img, ok := img.(*image.RGBA); ok {
		return img, nil
	}

	rect := image.Rectangle{image.Pt(0, 0), img.Bounds().Size()}
	dst := image.NewRGBA(rect)

	draw.Draw(dst, rect, img, image.Pt(0, 0), draw.Src)
	return dst, nil
}

func drawRect(img *image.RGBA, topLeftX, topLeftY, bottomRightX, bottomRightY int) {
	col := color.RGBA{255, 0, 0, 255}

	// horizontal
	for x := topLeftX; x < bottomRightX; x++ {
		// top
		img.Set(x, topLeftY, col)
		// bottom
		img.Set(x, bottomRightY, col)
	}

	// veritcal
	for y := topLeftY; y < bottomRightY; y++ {
		// left
		img.Set(topLeftX, y, col)
		// right
		img.Set(bottomRightX, y, col)
	}
}

func main() {
	img, err := getRGBA("./abetomo.png")
	if err != nil {
		panic(err)
	}
	drawRect(img, 150, 150, 300, 300)
	outfile, err := os.Create("./out.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}
