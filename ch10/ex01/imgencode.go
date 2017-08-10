package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var kind = flag.String("o", "jpeg", "an output encoding likes jpeg, png, or gif.")

func main() {
	flag.Parse()
	if err := encode(os.Stdin, os.Stdout, *kind); err != nil {
		fmt.Fprintf(os.Stderr, "imgencode: %v\n", err)
		os.Exit(1)
	}
}

func encode(in io.Reader, out io.Writer, toKind string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch toKind {
	case "jpeg", "jpg", "JPEG", "JPG":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png", "PNG":
		return png.Encode(out, img)
	case "gif", "GIF":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	default:
		return errors.New("unknown encoding:" + toKind)
	}
}
