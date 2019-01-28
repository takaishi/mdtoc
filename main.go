package main

import (
	"fmt"
	"github.com/takaishi/mdtoc/mdtoc"
	"github.com/urfave/cli"
	"io/ioutil"

	"log"
	"os"
)

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
	mdtoc := mdtoc.MDToc{File: c.String("file"), InFile: c.Bool("in-file"), Level: c.Int("level")}

	input, err := ioutil.ReadFile(mdtoc.File)
	if err != nil {
		return err
	}

	toc := mdtoc.GenerateTOC(input)

	output, err := mdtoc.InsertTOC(string(input), toc)
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
