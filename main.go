package main

import (
	"errors"
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
		os.Exit(1)
	}

	toc := generateTOC(input)

	outputWithTOC(string(input), toc)
}

func outputWithTOC(input string, toc string) error {
	tocPos := strings.Index(input, "<!-- toc -->")
	if tocPos == -1 {
		return errors.New("Can not find toc_pos comment `<!-- toc -->`.")
	}
	tocStartPos := strings.Index(string(input), "<!-- toc:start -->")
	if tocStartPos != -1 {
		tocEndPos := strings.Index(string(input), "<!-- toc:end -->")
		if tocEndPos == -1 {
			return errors.New("Can not find toc end position comment `<!-- toc:end -->`.")
		}

		spos := tocPos + 12
		epos := tocEndPos + 16
		output := input[:spos] + "\n<!--toc:start -->\n" + toc + "\n<!-- toc:end -->" + input[epos:]
		fmt.Println(output)
	} else {

		spos := tocPos + 12
		epos := tocPos + 12
		output := input[:spos] + "\n<!-- toc:start -->\n" + toc + "\n<!-- toc:end -->\n" + input[epos:]
		fmt.Println(output)
	}
	return nil
}

func generateTOC(input []byte) string {
	parser := blackfriday.New()
	toc := ""
	node := parser.Parse(input)
	node.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if node.Type == blackfriday.Heading {
			anchor := string(node.FirstChild.Literal)
			anchor = strings.Replace(anchor, " ", "", -1)
			anchor = strings.Replace(anchor, ".", "", -1)

			toc = fmt.Sprintf("%s\n%s [%s](%s)", toc, strings.Repeat("*", node.Level), node.FirstChild.Literal, anchor)

			if node.Next != nil {
				*node = *node.Next
			} else {
				return blackfriday.Terminate
			}
		}
		return blackfriday.GoToNext
	})
	return toc
}
