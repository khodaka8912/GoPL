package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(returnLessEcho(arg))
	}
}

func returnLessEcho(s string) (ret string) {
	defer func() {
		recover()
		ret = s
	}()
	panic("there is no return!")
}
