package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

var max int

func newton(z complex128) color.Color {
	const iterations = 100
	v := z
	for n := 0; n < iterations; n++ {
		v = v - (v*v*v*v-1)/(4*v*v*v)
		if cmplx.Abs(v*v) == 1 {
			c := uint8(255 - math.Log10(float64(n))*255/2)
			switch {
			case real(v) == 1:
				return color.RGBA{c, 0, 0, 255}
			case real(v) == -1:
				return color.RGBA{0, c, 0, 255}
			case imag(v) == 1:
				return color.RGBA{c, c, 0, 255}
			case imag(v) == -1:
				return color.RGBA{0, 0, c, 255}

			}
		}
	}
	return color.Black
}
