package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/urfave/cli/v2"
// )

// func main() {
// 	app := &cli.App{
// 		Name:  "week4",
// 		Usage: "A simple CLI tool to demonstrate urfave/cli",
// 		Commands: []*cli.Command{
// 			{
// 				Name: "hello",
// 				// Aliases: []string{"h"}, // 去掉别名 h 以避免冲突
// 				Usage: "say hello to the world or a specific person",
// 				Action: func(c *cli.Context) error {
// 					name := c.String("name")
// 					if name == "" {
// 						fmt.Println("Hello, World!")
// 					} else {
// 						fmt.Printf("Hello, %s!\n", name)
// 					}
// 					return nil
// 				},
// 				Flags: []cli.Flag{
// 					&cli.StringFlag{
// 						Name:  "name",
// 						Usage: "the name of the person to greet",
// 					},
// 				},
// 			},
// 			{
// 				Name:    "version",
// 				Aliases: []string{"v"},
// 				Usage:   "print the version of the CLI tool",
// 				Action: func(c *cli.Context) error {
// 					fmt.Println("week4 v1.0.0")
// 					return nil
// 				},
// 			},
// 		},
// 	}

// 	err := app.Run(os.Args)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }

import (
	"week4/lesson01"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Parent struct {
// 	ID   int `gorm:"primary_key"`
// 	Name string
// }

// type Child struct {
// 	Parent
// 	Age int
// }

// func InitDB(dst ...interface{}) *gorm.DB {
// 	db, err := gorm.Open(mysql.Open("root:080657@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.AutoMigrate(dst...)

// 	return db
// }

func main() {
	db, err := gorm.Open(mysql.Open("root:080657@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	lesson01.Run(db)
	// lesson02.Run(db)
	// lesson03.Run(db)
	// lesson03_02.Run(db)
	// lesson03_03.Run(db)
	// lesson03_04.Run(db)
	// lesson04.Run(db)

	//InitDB(&Parent{}, &Child{})
}
