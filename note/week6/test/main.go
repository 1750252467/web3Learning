package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"context"

	"github.com/urfave/cli/v3"
)

func main() {
	// 使用正确的方式创建CLI应用程序和所有命令
	app := &cli.Command{
		Name:    "calculator",
		Usage:   "A simple calculator CLI application",
		Version: "1.0.0",
		// 直接在结构体中定义子命令
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Add two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					args := cmd.Args().Slice()
					if len(args) != 2 {
						return fmt.Errorf("add command requires exactly 2 arguments")
					}
					num1, num2 := parseFloat(args[0]), parseFloat(args[1])
					result := num1 + num2
					fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result)
					return nil
				},
			},
			{
				Name:  "subtract",
				Usage: "Subtract two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					args := cmd.Args().Slice()
					if len(args) != 2 {
						return fmt.Errorf("subtract command requires exactly 2 arguments")
					}
					num1, num2 := parseFloat(args[0]), parseFloat(args[1])
					result := num1 - num2
					fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
					return nil
				},
			},
			{
				Name:  "multiply",
				Usage: "Multiply two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					args := cmd.Args().Slice()
					if len(args) != 2 {
						return fmt.Errorf("multiply command requires exactly 2 arguments")
					}
					num1, num2 := parseFloat(args[0]), parseFloat(args[1])
					result := num1 * num2
					fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
					return nil
				},
			},
			{
				Name:  "divide",
				Usage: "Divide two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					args := cmd.Args().Slice()
					if len(args) != 2 {
						return fmt.Errorf("divide command requires exactly 2 arguments")
					}
					num1, num2 := parseFloat(args[0]), parseFloat(args[1])
					if num2 == 0 {
						return fmt.Errorf("cannot divide by zero")
					}
					result := num1 / num2
					fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
					return nil
				},
			},
			{
				Name:  "modulo",
				Usage: "Calculate modulo of two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					args := cmd.Args().Slice()
					if len(args) != 2 {
						return fmt.Errorf("modulo command requires exactly 2 arguments")
					}
					num1, num2 := parseFloat(args[0]), parseFloat(args[1])
					if num2 == 0 {
						return fmt.Errorf("cannot perform modulo with zero")
					}
					// Convert to integers for modulo operation
					intNum1, intNum2 := int(num1), int(num2)
					result := intNum1 % intNum2
					fmt.Printf("%d %% %d = %d\n", intNum1, intNum2, result)
					return nil
				},
			},
		},
	}

	// 运行应用程序
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

// Helper function to parse string to float64
func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Warning: Could not parse '%s' as number, using 0\n", s)
		return 0
	}
	return f
}
