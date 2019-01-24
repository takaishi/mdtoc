package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
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

	tocPos := strings.Index(string(input), "<!-- toc -->")
	if tocPos == -1 {
		log.Printf("[ERROR] can not find toc_pos comment `<!-- toc -->`.")
		os.Exit(1)
	}
	tocStartPos := strings.Index(string(input), "<!-- toc:start -->")
	if tocStartPos != -1 {
		tocEndPos := strings.Index(string(input), "<!-- toc:end -->")
		if tocEndPos == -1 {
			log.Printf("[ERROR] can not find toc end position comment `<!-- toc:end -->`.")
			os.Exit(1)
		}

		s := string(input)
		spos := tocPos + 12
		epos := tocEndPos + 16
		output := s[:spos] + "\n<!--toc:start -->\n" + toc + "\n<!-- toc:end -->" + s[epos:]
		fmt.Println(output)
	} else {

		pos := tocPos + 12
		s := string(input)
		output := s[:pos] + "\n<!-- toc:start -->\n" + toc + "\n<!-- toc:end -->\n" + s[pos:]
		fmt.Println(output)

	}

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
