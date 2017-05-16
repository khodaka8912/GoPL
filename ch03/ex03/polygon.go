package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 400
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.3
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var fn = f

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "saddle" {
			fn = saddle
		} else if args[0] == "egg" {
			fn = egg
		}
	}
	min, max := minmax()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aInf := corner(i+1, j)
			bx, by, bInf := corner(i, j)
			cx, cy, cInf := corner(i, j+1)
			dx, dy, dInf := corner(i+1, j+1)
			if aInf || bInf || cInf || dInf {
				continue
			}
			color := color(i, j, min, max)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := fn(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	sInf := math.IsInf(z, 0) || math.IsNaN(z)
	return sx, sy, sInf
}

func minmax() (float64, float64) {
	min, max := math.Inf(1), math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)

			z := fn(x, y)

			if math.IsInf(z, 0) || math.IsNaN(z) {
				continue
			}
			min, max = math.Min(min, z), math.Max(max, z)
		}
	}
	return min, max
}

func color(i, j int, min, max float64) string {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := fn(x, y)

	half := (max - min) / 2
	center := (max + min) / 2
	if z > center {
		red := uint8(255 * (z - center) / half)
		return fmt.Sprintf("#%02x%02x00", red, 255-red)
	} else {
		blue := uint8(255 * (center - z) / half)
		return fmt.Sprintf("#00%02x%02x", 255-blue, blue)
	}

}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func saddle(x, y float64) float64 {
	return (x*x/3-y*y)/200 + 0.5
}

func egg(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 5
}
