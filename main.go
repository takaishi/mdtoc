package main

import (
	"errors"
	"fmt"
	"github.com/russross/blackfriday"
	"github.com/urfave/cli"
	"io/ioutil"

	"log"
	"os"
	"strings"
)

const TOC_POS = "<!-- toc -->"
const TOC_START_POS = "<!-- toc:start -->"
const TOC_END_POS = "<!-- toc:end -->"

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "file, f",
		},
	}

	app.Action = func(c *cli.Context) error {
		return action(c)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func action(c *cli.Context) error {
	input, err := ioutil.ReadFile(c.String("file"))
	if err != nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}

	toc := generateTOC(input)

	outputWithTOC(string(input), toc)
	return nil
}

func outputWithTOC(input string, toc string) error {
	tocPos := strings.Index(input, TOC_POS)
	if tocPos == -1 {
		return errors.New(fmt.Sprintf("Can not find toc_pos comment `%s`", TOC_POS))
	}
	tocStartPos := strings.Index(string(input), TOC_START_POS)
	if tocStartPos != -1 {
		tocEndPos := strings.Index(string(input), TOC_END_POS)
		if tocEndPos == -1 {
			return errors.New(fmt.Sprintf("Can not find toc end position comment `%s`.", TOC_END_POS))
		}

		spos := tocPos + 12
		epos := tocEndPos + 16
		output := input[:spos] + "\n" + TOC_START_POS + "\n" + toc + "\n" + TOC_END_POS + input[epos:]
		fmt.Printf(output)
	} else {

		spos := tocPos + 12
		epos := tocPos + 12
		output := input[:spos] + "\n" + TOC_START_POS + "\n" + toc + "\n" + TOC_END_POS + input[epos:]
		fmt.Printf(output)
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
