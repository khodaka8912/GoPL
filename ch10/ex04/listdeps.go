package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	arg := os.Args[1]
	out, err := exec.Command("go", "list", "-e", "-json", arg).Output()
	decoder := json.NewDecoder(bytes.NewReader(out))
	target := Package{}
	err = decoder.Decode(&target)
	if err != nil && err != io.EOF {
		log.Fatalf("go list err:%v\n", err)
	}

	out, err = exec.Command("go", "list", "-e", "-json", "...").Output()
	if err != nil {
		log.Fatalf("go list err:%v\n", err)
	}
	decoder = json.NewDecoder(bytes.NewReader(out))
	packages := []Package{}
	for {
		pack := Package{}
		err := decoder.Decode(&pack)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "json decode err:%v\n", err)
			continue
		}
		for _, dep := range pack.Deps {
			if dep == target.ImportPath {
				packages = append(packages, pack)
				break
			}
		}
	}
	fmt.Printf("%s (%s) is depended by:\n", target.Name, target.ImportPath)
	for _, pack := range packages {
		fmt.Printf("\t%s (%s)\n", pack.Name, pack.ImportPath)
	}
}

type Package struct {
	ImportPath string   // import path of package in dir
	Name       string   // package name
	Deps       []string // all (recursively) imported dependencies
}
