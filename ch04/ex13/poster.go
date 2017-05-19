package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var title = flag.String("title", "", "search for title")

func main() {
	flag.Parse()
	movie := findMovie(*title)
	if movie.Response == "False" {
		fmt.Println(movie.Error)
		return
	}
	fmt.Printf("found: %s\n", movie.Title)
	downloadPoster(movie.Poster)
}

type movie struct {
	Title    string
	Poster   string
	Response string
	Error    string
}

func findMovie(title string) movie {
	url := fmt.Sprintf("http://www.omdbapi.com/?t=%s", url.QueryEscape(title))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	movie := movie{}
	if err = decoder.Decode(&movie); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return movie
}

func downloadPoster(poster string) {
	if poster == "N/A" {
		fmt.Println("poster not available")
		return
	}
	resp, err := http.Get(poster)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	filename := poster[strings.LastIndex(poster, "/")+1:]
	err = ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("downloaded: %s\n", filename)
}
