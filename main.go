package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./test.md")
	if err != nil {
		fmt.Println("failed to read file")
		os.Exit(-1)
	}
	parser := blackfriday.New()
	node := parser.Parse(input)

	node.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if node.Type == blackfriday.Heading {
			anchor := string(node.FirstChild.Literal)
			anchor = strings.Replace(anchor, " ", "", -1)
			anchor = strings.Replace(anchor, ".", "", -1)

			fmt.Printf("%s [%s](%s)\n", strings.Repeat("*", node.Level), node.FirstChild.Literal, anchor)

			if node.Next != nil {
				*node = *node.Next
			} else {
				return blackfriday.Terminate
			}
		}
		return blackfriday.GoToNext
	})
}
