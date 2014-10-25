package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	// TODO: Directory here driven from a config object
	app := cli.NewApp()
	app.Name = "lanyon"
	app.Usage = "A JSON based static site generator"
	app.Commands = []cli.Command{
		{
			Name:      "server",
			ShortName: "s",
			Usage:     "start a lanyon server",
			Action: func(c *cli.Context) {
				Server()
			},
		},
		{
			Name:      "build",
			ShortName: "b",
			Usage:     "compile site from static JSON",
			Action: func(c *cli.Context) {
				NewBlog("_posts/")
			},
		},
	}
	app.Run(os.Args)
}
