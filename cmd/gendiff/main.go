package main

import (
	"code"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		UsageText: "gendiff [global options] <filepath1> <filepath2>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Value:   "stylish",
				Usage:   "output format",
				Aliases: []string{"f"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			filepath1 := cmd.Args().Get(0)
			if filepath1 == "" {
				return errors.New("path1 is required")
			}

			filepath2 := cmd.Args().Get(1)
			if filepath2 == "" {
				return errors.New("path2 is required")
			}

			result, err := code.GenDiff(filepath1, filepath2, "stylish")
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
