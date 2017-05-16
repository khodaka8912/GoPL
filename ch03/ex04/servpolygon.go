package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const (
	width, height = 600, 400
	cells         = 100
	xyrange       = 30.0
	angle         = math.Pi / 6
	grad          = "gradation"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type query struct {
	width  int
	height int
	color  string
	fn     func(float64, float64) float64
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		q := getQuery(r.Form)
		w.Header().Set("Content-Type", "image/svg+xml")
		polygon(w, q)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getQuery(form url.Values) query {
	q := query{width, height, grad, f}
	fmt.Println(form)
	if v := form.Get("width"); len(v) != 0 {
		w, err := strconv.Atoi(v)
		if err != nil {
			log.Print(err)
		} else {
			q.width = w
		}
	}
	if v := form.Get("height"); len(v) != 0 {
		h, err := strconv.Atoi(v)
		if err != nil {
			log.Print(err)
		} else {
			q.height = h
		}
	}
	if v := form.Get("color"); len(v) != 0 {
		if matched, err := regexp.MatchString(`^(#[0-9a-fA-F]{6}|gradeation)$`, v); err != nil {
			log.Print(err)
		} else {
			if matched {
				q.color = v
			} else {
				log.Printf("unknown color:%q", v)
			}
		}
	}
	if v := form.Get("func"); len(v) != 0 {
		switch v {
		case "egg":
			q.fn = egg
		case "saddle":
			q.fn = saddle
		default:
			log.Printf("unknown func:%q", v)
		}
	}
	return q
}

func polygon(out io.Writer, q query) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", q.width, q.height)
	min, max := minmax(q)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aInf := corner(i+1, j, q)
			bx, by, bInf := corner(i, j, q)
			cx, cy, cInf := corner(i, j+1, q)
			dx, dy, dInf := corner(i+1, j+1, q)
			if aInf || bInf || cInf || dInf {
				continue
			}
			fillColor := q.color
			if q.color == grad {
				fillColor = color(i, j, min, max, q)

			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fillColor)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, q query) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	xyscale := float64(q.width / 2 / xyrange)
	zscale := float64(q.height) * 0.3

	z := q.fn(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	sInf := math.IsInf(sx, 0) || math.IsInf(sy, 0)
	return sx, sy, sInf
}

func minmax(q query) (float64, float64) {
	min, max := math.Inf(1), math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)

			z := q.fn(x, y)

			if math.IsInf(z, 0) || math.IsNaN(z) {
				continue
			}
			min, max = math.Min(min, z), math.Max(max, z)
		}
	}
	return min, max
}

func color(i, j int, min, max float64, q query) string {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := q.fn(x, y)

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
