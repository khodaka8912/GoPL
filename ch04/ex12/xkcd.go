package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var title = flag.String("title", "", "search for title")

func main() {
	flag.Parse()
	if len(*title) > 0 {
		search(*title)
		return
	}
	arg := flag.Args()
	if len(arg) != 1 {
		fmt.Fprintln(os.Stderr, "xkcd [-serch word] [comic number]")
		os.Exit(1)
	}
	n, err := strconv.Atoi(arg[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "xkcd [-serch word] [comic number]")
		os.Exit(1)
	}
	download(n)
}

func download(n int) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	filename := fmt.Sprintf("%d.json", n)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	ioutil.WriteFile(filename, data, os.ModePerm)
}

type info0 struct {
	Num        int
	Link       string
	Title      string
	SafeTitle  string `json:"safe_title"`
	Year       string
	Month      string
	Day        string
	Transcript string
	Img        string
	Alt        string
	News       string
}

func search(title string) {
	files, err := filepath.Glob("*.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
			continue
		}
		info := info0{}
		err = json.Unmarshal(data, &info)
		if err != nil {
			log.Fatal(err)
			continue
		}
		if strings.Contains(info.Title, title) {
			fmt.Printf("https://xkcd.com/%d/\n", info.Num)
			fmt.Println(info.Transcript)
		}
	}
}
