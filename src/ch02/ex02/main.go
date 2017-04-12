package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			printUnitConv(arg, os.Stdout)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			printUnitConv(input.Text(), os.Stdout)
		}
	}
}

func printUnitConv(s string, out *os.File) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conv: %v\n", err)
	}
	const format = "%v = %v\n"
	c := Celsius(value)
	f := Fahrenheit(value)
	fmt.Fprintf(out, format, c, CToF(c))
	fmt.Fprintf(out, format, f, FToC(f))
	m := Meter(value)
	ft := Feet(value)
	fmt.Fprintf(out, format, m, MToFt(m))
	fmt.Fprintf(out, format, ft, FtToM(ft))
	kg := Kilogram(value)
	lb := Pound(value)
	fmt.Fprintf(out, format, kg, KgToLb(kg))
	fmt.Fprintf(out, format, lb, LbToKg(lb))
}
