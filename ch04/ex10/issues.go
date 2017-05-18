package main

import (
	"os"

	"fmt"
	"log"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	im, iy, oy := classifyTime(result.Items)
	fmt.Println("In a month:")
	for _, item := range im {
		fmt.Printf("%v #%-5d %9.9s %.55s\n", item.CreatedAt, item.Number, item.User.Login, item.Title)
	}
	fmt.Println("In a year:")
	for _, item := range iy {
		fmt.Printf("%v #%-5d %9.9s %.55s\n", item.CreatedAt, item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Over a year ago:")
	for _, item := range oy {
		fmt.Printf("%v #%-5d %9.9s %.55s\n", item.CreatedAt, item.Number, item.User.Login, item.Title)
	}
}

func classifyTime(items []*github.Issue) ([]*github.Issue, []*github.Issue, []*github.Issue) {
	const (
		monthAsHours = 24 * 30
		yearAsHours  = 24 * 365
	)
	inMonth, inYear, overYear := []*github.Issue{}, []*github.Issue{}, []*github.Issue{}
	now := time.Now()
	for _, item := range items {
		hours := now.Sub(item.CreatedAt).Hours()
		switch {
		case hours < monthAsHours:
			inMonth = append(inMonth, item)
		case hours < yearAsHours:
			inYear = append(inYear, item)
		default:
			overYear = append(overYear, item)
		}
	}
	return inMonth, inYear, overYear
}
