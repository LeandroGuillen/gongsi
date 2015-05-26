package main

import (
	"github.com/leandroguillen/gongsi"
	"github.com/codegangsta/cli"
	"os"
	"strconv"
	"fmt"
)

var configFile = "/tmp/ngsi.config"
var g gongsi.Gongsi

func main() {

	// Set up Gongsi
	g = gongsi.Gongsi{
		Encoding: "application/json",
		Host:     "localhost",
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
		fmt.Println("USAGE: ngsi [global options] command [command options] [arguments...]")
		fmt.Println("Type 'ngsi help' for help")
	}

	app.Commands = []cli.Command{
		{
			Name:      "query",
			ShortName: "q",
			Usage:     "Query operation",
			Action:    query,
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
		{
			Name:      "host",
			Usage:     "Set up context broker host",
			Action:    host,
		},
	}

	app.Run(os.Args)

}

func query(c *cli.Context) {
	arg1 := c.Args().Get(0)
	arg2 := c.Args().Get(1)

	if arg1 != "" && arg2 != "" {
		// Query for an attribute
		out, err := g.ConvQueryContextAttribute(arg1, arg2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(out)
		}
	} else if arg1 != "" {
		// Query for an entity
		out, err := g.ConvQueryContext(arg1)
		if err != nil {
			fmt.Println(err)	
		} else {
			fmt.Println(out)
		}
	} else {
		fmt.Println("You must provide an ID:")
		fmt.Println("   ngsi get <ID>")
	}		
}

func host(c *cli.Context) {				
	arg1 := c.Args().Get(0)
	arg2 := c.Args().Get(1)

	if arg1 != "" {
		
		// Setting up the host
		fmt.Println("setting up host")
		if _, err := os.Stat(configFile); err == nil {
		    fmt.Println(configFile + " exists; processing...")
		} else {
			fmt.Println("file does not exist")
		}
			
	} else if arg1 != "" && arg2 != "" {
		// Setting up the host and the port
		fmt.Println("setting up host and port")

	} else {
		// Get info
		fmt.Println("Current settings:")
		fmt.Println("  host: " + g.Host)
		fmt.Println("  port: " + strconv.Itoa(g.Port))
		

	}
		
}