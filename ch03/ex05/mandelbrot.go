package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const sep = 10

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			c := uint8(n % sep * 255 / (sep - 1))
			switch n / sep {
			case 0:
				return color.RGBA{255, c, 0, 255}
			case 1:
				return color.RGBA{255 - c, 255, 0, 255}
			case 2:
				return color.RGBA{0, 255, c, 255}
			case 3:
				return color.RGBA{0, 255 - c, 255, 255}
			case 4:
				return color.RGBA{c, 0, 255, 255}
			default:
				return color.RGBA{255, 0, 255, 255}

			}
		}
	}
	return color.Black
}
