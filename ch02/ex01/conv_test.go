package tempconv

import (
	"math"
	"testing"
)

type TempSet struct {
	c Celsius
	f Fahrenheit
	k Kelvin
}

var testCases = []TempSet{
	{0, 32, 273.15},
	{55.55, 131.99, 328.7},
	{-12, 10.4, 261.15},
}

func TestCToF(t *testing.T) {
	for _, temp := range testCases {
		f := CToF(temp.c)
		if math.Abs(float64(f-temp.f)) > 0.00001 {
			t.Errorf("CToF(%q) want %q but %q", temp.c, temp.f, f)
		}
	}
}

func TestCToK(t *testing.T) {
	for _, temp := range testCases {
		k := CToK(temp.c)
		if math.Abs(float64(k-temp.k)) > 0.00001 {
			t.Errorf("CToK(%q) want %q but %q", temp.c, temp.k, k)
		}
	}
}

func TestFToC(t *testing.T) {
	for _, temp := range testCases {
		c := FToC(temp.f)
		if math.Abs(float64(c-temp.c)) > 0.00001 {
			t.Errorf("FToC(%q) want %q but %q", temp.f, temp.c, c)
		}
	}
}

func TestFToK(t *testing.T) {
	for _, temp := range testCases {
		k := FToK(temp.f)
		if math.Abs(float64(k-temp.k)) > 0.00001 {
			t.Errorf("FToK(%q) want %q but %q", temp.f, temp.k, k)
		}
	}
}

func TestKToC(t *testing.T) {
	for _, temp := range testCases {
		c := KToC(temp.k)
		if math.Abs(float64(c-temp.c)) > 0.00001 {
			t.Errorf("KToC(%q) want %q but %q", temp.k, temp.c, c)
		}
	}
}

func TestKToF(t *testing.T) {
	for _, temp := range testCases {
		f := KToF(temp.k)
		if math.Abs(float64(f-temp.f)) > 0.00001 {
			t.Errorf("KToF(%q) want %q but %q", temp.k, temp.f, f)
		}
	}
}
