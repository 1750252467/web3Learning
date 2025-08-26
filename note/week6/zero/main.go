package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			name := "world"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}

			if c.String("lang") == "english" {
				fmt.Println("hello", name)
			} else {
				fmt.Println("你好", name)
			}
			return nil
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}
