package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

var useSha384 = flag.Bool("sha384", false, "create hash of sha384")
var useSha512 = flag.Bool("sha512", false, "create hash of sha512")

func main() {
	flag.Parse()
	if *useSha384 && *useSha512 {
		fmt.Errorf("The flags -sha384 and -sha512 cannot be used at the same time.\n")
		return
	}
	var hash hash.Hash
	if *useSha384 {
		hash = sha512.New384()
	} else if *useSha512 {
		hash = sha512.New()
	} else {
		hash = sha256.New()
	}
	sum := sum(os.Stdin, hash)
	fmt.Println(sum)
}

func sum(in io.Reader, hash hash.Hash) string {
	if _, err := io.Copy(hash, in); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
