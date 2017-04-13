package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	blackIndex = 0
	greenIndex = 1
)

const (
	cycles  = 5
	res     = 0.001
	size    = 100
	nframes = 64
	delay   = 8
)

type query struct {
	cycles  int
	size    int
	nframes int
	delay   int
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		q := getQuery(r.Form)
		lissajous(w, q)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getQuery(form url.Values) query {
	q := query{cycles, size, nframes, delay}
	if v := form.Get("cycles"); len(v) != 0 {
		c, err := strconv.Atoi(v)
		if err != nil {
			log.Print(err)
		} else {
			q.cycles = c
		}
	}
	if v := form.Get("size"); len(v) != 0 {
		s, err := strconv.Atoi(v)
		if err != nil {
			log.Print(err)
		} else {
			q.size = s
		}
	}
	if v := form.Get("nframes"); len(v) != 0 {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Print(err)
		} else {
			q.nframes = n
		}
	}
	if v := form.Get("delay"); len(v) != 0 {
		d, err := strconv.Atoi(v)
		if err != nil {
			log.Print(err)
		} else {
			q.delay = d
		}
	}
	return q
}

func lissajous(out io.Writer, q query) {
	fmt.Printf("query:%v\n", q)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: q.nframes}
	phase := 0.0
	for i := 0; i < q.nframes; i++ {
		rect := image.Rect(0, 0, 2*q.size+1, 2*q.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(q.cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(q.size+int(x*float64(q.size)+0.5), q.size+int(y*float64(q.size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, q.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
