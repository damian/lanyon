package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var config, err = NewConfig("config.json")

func main() {
	app := cli.NewApp()
	app.Name = "lanyon"
	app.Usage = "A JSON based static site generator"
	app.Commands = []cli.Command{
		{
			Name:      "server",
			ShortName: "s",
			Usage:     "start a lanyon server",
			Action: func(c *cli.Context) {
				NewBlog(config.Source)
				Server()
			},
		},
		{
			Name:      "build",
			ShortName: "b",
			Usage:     "compile site from static JSON",
			Action: func(c *cli.Context) {
				NewBlog(config.Source)
			},
		},
		{
			Name:  "new",
			Usage: "scaffolds a new lanyon website",
			Action: func(c *cli.Context) {
				println("Generated:", c.Args().First())
			},
		},
	}
	app.Run(os.Args)
}
