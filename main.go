package main

import (
	"fmt"
	"os"

	"github.com/mh-cbon/go-repo-utils/repoutils"
	"github.com/urfave/cli"
)

var VERSION = "0.0.0"

func main() {

	app := cli.NewApp()
	app.Name = "commit"
	app.Version = VERSION
	app.Usage = "Commit file"
	app.UsageText = "commit -m <message> -f <file>"
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "file, f",
			Value: &cli.StringSlice{},
			Usage: "File to add and commit",
		},
		cli.StringFlag{
			Name:  "message, m",
			Value: "",
			Usage: "Message of the commit",
		},
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "Silently fail",
		},
	}

	app.Action = func(c *cli.Context) error {
		files := c.StringSlice("file")
		message := c.String("message")
		quiet := c.Bool("quiet")

		if len(files) == 0 {
			cli.ShowAppHelp(c)
			return cli.NewExitError("Files are required", 1)
		}

		if len(message) == 0 {
			cli.ShowAppHelp(c)
			return cli.NewExitError("Message is required", 1)
		}

		path, err := os.Getwd()
		exitWithError(err)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		vcs, err := repoutils.WhichVcs(path)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		sfile := make([]string, 0)
		for _, file := range files {
			sfile = append(sfile, string(file))
			err = repoutils.Add(vcs, path, string(file))
			if err != nil {
				if quiet {
					// it does not exit on error, just print it
					fmt.Printf("Failed to add %s: %s\n", file, err.Error())
				} else {
					return cli.NewExitError(err.Error(), 1)
				}
			}
		}

		err = repoutils.Commit(vcs, path, message, sfile)
		if err != nil {
			if quiet {
				// it does not exit on error, just print it
				fmt.Printf("Failed to commit %s\n", err.Error())
			} else {
				return cli.NewExitError(err.Error(), 1)
			}
		}

		return nil
	}

	app.Run(os.Args)
}

func exitWithError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
