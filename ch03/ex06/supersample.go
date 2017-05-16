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
	dx, dy := float64(xmax-xmin)/width/4, float64(ymax-ymin)/height/4
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, supersample(x, y, dx, dy))
		}
	}
	png.Encode(os.Stdout, img)
}

func supersample(x, y, dx, dy float64) color.Color {
	c1 := mandelbrot(complex(x+dx, y+dy))
	c2 := mandelbrot(complex(x+dx, y-dy))
	c3 := mandelbrot(complex(x-dx, y+dy))
	c4 := mandelbrot(complex(x-dx, y-dy))
	return average(c1, c2, c3, c4)
}

func average(colors ...color.RGBA) color.RGBA {
	size := uint32(len(colors))
	var r, g, b, a uint32
	for _, c := range colors {
		r += uint32(c.R)
		g += uint32(c.G)
		b += uint32(c.B)
		a += uint32(c.A)
	}
	return color.RGBA{uint8(r / size), uint8(g / size), uint8(b / size), uint8(a / size)}
}

func mandelbrot(z complex128) color.RGBA {
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
	return color.RGBA{0, 0, 0, 255}
}
