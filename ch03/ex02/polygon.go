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

// go run polygon.go saddle
// go run polygon.go egg
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	fn := saddle
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "saddle" {
			fn = saddle
		} else if args[0] == "egg" {
			fn = egg
		}
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aInf := corner(i+1, j, fn)
			bx, by, bInf := corner(i, j, fn)
			cx, cy, cInf := corner(i, j+1, fn)
			dx, dy, dInf := corner(i+1, j+1, fn)
			if aInf || bInf || cInf || dInf {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f func(float64, float64) float64) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	sInf := math.IsInf(sx, 0) || math.IsInf(sy, 0)
	return sx, sy, sInf
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func saddle(x, y float64) float64 {
	return (x*x/3-y*y) / 200 + 0.5
}

func egg(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 5
}