package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var config, err = NewConfig("config.json")
var blog *Blog

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
				CopyStaticAssets(config.Static, config.Destination)
				blog, _ = NewBlog(config.Source)
				blog.save()
				Server()
			},
		},
		{
			Name:      "build",
			ShortName: "b",
			Usage:     "compile site from static JSON",
			Action: func(c *cli.Context) {
				CopyStaticAssets(config.Static, config.Destination)
				blog, _ = NewBlog(config.Source)
				blog.save()
			},
		},
		{
			Name:  "editor",
			Usage: "launches editor mode for writing new blog posts",
			Action: func(c *cli.Context) {
				blog, _ = NewBlog(config.Source)
				NewEditor()
			},
		},
		{
			Name:  "new",
			Usage: "scaffolds a new lanyon website",
			Action: func(c *cli.Context) {
				// TODO: Build a lanyon scaffolder
				println("Generated:", c.Args().First())
			},
		},
	}
	app.Run(os.Args)
}
