package main

import "fmt"

type Meter float64
type Feet float64

const (
	OneFoot Meter = 0.3048
)

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func (ft Feet) String() string { return fmt.Sprintf("%gft", ft) }

func MToFt(m Meter) Feet { return Feet(m / OneFoot) }

func FtToM(ft Feet) Meter { return Meter(ft) * OneFoot }
