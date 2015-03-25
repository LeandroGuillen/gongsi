package main

import (
	"github.com/leandroguillen/gongsi"
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	// Set up Gongsi
	g := gongsi.Gongsi{
		Encoding: "application/json",
		Host:     "leandroguillen.com",
		Port:     1026,
	}
	g.Init()

	app := cli.NewApp()
	app.Name = "ngsi"
	app.Usage = "NGSI command-line helper"
	app.Version = "0.1"
	app.Author = "Leandro Guillen"
	app.Email = "leandro.guillen@imdea.org"

	// Default operation with no parameters
	// app.Action = func(c *cli.Context) {
	// 	println("I ngsi U")
	// }

	app.Commands = []cli.Command{
		{
			Name:      "query",
			ShortName: "q",
			Usage:     "Query operation",
			Action: func(c *cli.Context) {
				if str := c.Args().First(); str != "" {
					// println("Query an entity by ID: ", str)
					out, _ := g.ConvQueryContext(str)
					println(out)
				} else {
					println("You must provide an ID:")
					println("   ngsi get <ID>")
				}
			},
		},
		{
			Name:      "update",
			ShortName: "u",
			Usage:     "Update operation",
			Action: func(c *cli.Context) {
				println("UPDATE operation", c.Args().First())
			},
		},
		{
			Name:      "subscribe",
			ShortName: "s",
			Usage:     "Subscribe operation",
			Action: func(c *cli.Context) {
				println("SUBSCRIBE operation", c.Args().First())
			},
		},
	}

	app.Run(os.Args)

}
