package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	c, i := CountingWriter(os.Stdout)
	for _, arg := range os.Args[1:] {
		io.WriteString(c, arg)
	}
	fmt.Println()
	fmt.Printf("%d bytes are written\n", *i)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &countingWriter{w: w}
	return c, &c.i
}

type countingWriter struct {
	w io.Writer
	i int64
}

func (c *countingWriter) Write(p []byte) (int, error) {
	size, err := c.w.Write(p)
	c.i += int64(size)
	return size, err
}
