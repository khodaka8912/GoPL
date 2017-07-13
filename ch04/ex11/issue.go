package main

import (
	"fmt"
	"log"

	"os"

	"gopl.io/ch4/github"
)

var repo = []string{"repo:khodaka8912/Issue", "is:open"}

func main() {
	if len(os.Args[1:]) == 0 {
		usage()
		return
	}
	command := os.Args[1]
	args := os.Args[2:]
	switch command {
	case "help":
		usage()
	case "list":
		list(args)
	case "show":
		show(args)
	case "create":
		create(args)
	case "update":
		update(args)
	case "close":
		close(args)
	default:
		usage()
	}

}

func usage() {
	const text = `usage:
	issue [command] [options]...
commands are:
	list	show a list of open issues
	show	show information of a issue
	create	create a new issue
	update	update an existing issue
	close	close an existing issue
	help	show this usage message`
	fmt.Println(text)
}

func list(args []string) {
	result, err := github.SearchIssues(repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("%v #%-5d %9.9s %.55s\n", item.CreatedAt, item.Number, item.User.Login, item.Title)
	}
}

func create(args []string) {
	fmt.Println("create")
}

func show(args []string) {
	fmt.Println("show")

}

func update(args []string) {
	fmt.Println("update")
}

func close(args []string) {
	fmt.Println("close")
}
