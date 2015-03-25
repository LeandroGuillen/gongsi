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
	app.Action = func(c *cli.Context) {
		println("USAGE: ngsi [global options] command [command options] [arguments...]")
		println("Type 'ngsi help' for help")
	}

	app.Commands = []cli.Command{
		{
			Name:      "query",
			ShortName: "q",
			Usage:     "Query operation",
			Action: func(c *cli.Context) {
				
				arg1 := c.Args().Get(0)
				arg2 := c.Args().Get(1)

				if arg1 != "" && arg2 != "" {
					// Query for an attribute
					out, _ := g.ConvQueryContextAttribute(arg1, arg2)
					println(out)
				} else if arg1 != "" {
					// Query for an entity
					out, _ := g.ConvQueryContext(arg1)
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
