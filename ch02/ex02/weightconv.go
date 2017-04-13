package main

import "fmt"

type Kilogram float64
type Pound float64

const (
	OnePound Kilogram = 0.45359237
)

func (kg Kilogram) String() string { return fmt.Sprintf("%fkg", kg) }

func (lb Pound) String() string { return fmt.Sprintf("%flb", lb) }

func KgToLb(kg Kilogram) Pound { return Pound(kg / OnePound) }

func LbToKg(lb Pound) Kilogram { return Kilogram(lb) * OnePound }
