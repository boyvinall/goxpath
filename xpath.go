package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
)

func parse(c *cli.Context, data []byte) {
	doc, err := xml.Parse(data, nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	defer doc.Free()

	xp := doc.DocXPathCtx()
	for _, xmlns := range c.StringSlice("xmlns") {
		ns := strings.SplitN(xmlns, ":", 2)
		if c.Bool("verbose") {
			fmt.Println("NS " + ns[0] + "==" + ns[1])
		}
		xp.RegisterNamespace(ns[0], ns[1])
	}

	xps := xpath.Compile(c.String("xpath"))
	s, err := doc.Root().Search(xps)
	util.CheckErr(err)
	for _, s := range s {
		util.CheckErr(err)
		fmt.Println(s.Content())
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "xpath"
	app.Usage = "print result of xpath query"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "file",
			Value:  "-",
			Usage:  "input filename (default stdin)",
			EnvVar: "XP_FILE",
		},
		cli.StringFlag{
			Name:   "xpath",
			Value:  ".",
			Usage:  "xpath query",
			EnvVar: "XP_XPATH",
		},
		cli.BoolFlag{
			Name:   "verbose,V",
			Usage:  "print verbose output",
			EnvVar: "XP_VERBOSE",
		},
		cli.StringSliceFlag{
			Name:   "xmlns",
			Value:  &cli.StringSlice{"atom:http://www.w3.org/2005/Atom"},
			Usage:  "register XML namespaces using ns:uri",
			EnvVar: "XP_XMLNS",
		},
	}
	app.Action = func(c *cli.Context) {
		filename := c.String("file")
		var data []byte
		var err error
		if filename == "-" {
			data, err = ioutil.ReadAll(os.Stdin)
		} else {
			data, err = ioutil.ReadFile(filename)
		}
		util.CheckErr(err)
		parse(c, data)
	}

	app.Run(os.Args)
}
