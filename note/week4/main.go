package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "week4",
		Usage: "A simple CLI tool to demonstrate urfave/cli",
		Commands: []*cli.Command{
			{
				Name: "hello",
				// Aliases: []string{"h"}, // 去掉别名 h 以避免冲突
				Usage: "say hello to the world or a specific person",
				Action: func(c *cli.Context) error {
					name := c.String("name")
					if name == "" {
						fmt.Println("Hello, World!")
					} else {
						fmt.Printf("Hello, %s!\n", name)
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "the name of the person to greet",
					},
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "print the version of the CLI tool",
				Action: func(c *cli.Context) error {
					fmt.Println("week4 v1.0.0")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
