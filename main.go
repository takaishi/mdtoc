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

type MDToc struct {
	File   string
	InFile bool
	Level  int
}

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Specify to generate TOC.",
		},
		cli.BoolFlag{
			Name:  "in-file, i",
			Usage: "Insert TOC to md file specified --file option.",
		},
		cli.IntFlag{
			Name:  "level, l",
			Value: 2,
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
	mdtoc := MDToc{File: c.String("file"), InFile: c.Bool("in-file"), Level: c.Int("level")}

	input, err := ioutil.ReadFile(mdtoc.File)
	if err != nil {
		return err
	}

	toc := mdtoc.generateTOC(input)

	output, err := mdtoc.generateWithTOC(string(input), toc)
	if err != nil {
		return err
	}

	if c.Bool("in-file") {
		f, err := os.OpenFile(c.String("file"), os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
		defer f.Close()
		if err != nil {
			return err
		}

		_, err = f.Write([]byte(output))
		if err != nil {
			return err
		}
	} else {
		fmt.Printf(output)
	}
	return nil
}

func (mdtoc *MDToc) generateWithTOC(input string, toc string) (string, error) {
	tocPos := strings.Index(input, TOC_POS)
	if tocPos == -1 {
		return "", errors.New(fmt.Sprintf("Can not find toc_pos comment `%s`", TOC_POS))
	}
	tocStartPos := strings.Index(string(input), TOC_START_POS)
	if tocStartPos != -1 {
		tocEndPos := strings.Index(string(input), TOC_END_POS)
		if tocEndPos == -1 {
			return "", errors.New(fmt.Sprintf("Can not find toc end position comment `%s`.", TOC_END_POS))
		}

		spos := tocPos + 12
		epos := tocEndPos + 16
		output := input[:spos] + "\n" + TOC_START_POS + "\n" + toc + "\n\n" + TOC_END_POS + input[epos:]
		return output, nil
	} else {

		spos := tocPos + 12
		epos := tocPos + 12
		output := input[:spos] + "\n" + TOC_START_POS + "\n" + toc + "\n\n" + TOC_END_POS + input[epos:]
		return output, nil
	}
	return "", nil
}

func (mdtoc *MDToc) generateTOC(input []byte) string {
	parser := blackfriday.New()
	toc := ""
	node := parser.Parse(input)
	node.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if node.Type == blackfriday.Heading && node.Level > 1 && node.Level <= mdtoc.Level {
			anchor := string(node.FirstChild.Literal)
			anchor = strings.Replace(anchor, " ", "", -1)
			anchor = strings.Replace(anchor, ".", "", -1)
			anchor = strings.ToLower(anchor)

			toc = fmt.Sprintf("%s\n%s* [%s](#%s)", toc, strings.Repeat(" ", (node.Level-1)*2), node.FirstChild.Literal, anchor)

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
