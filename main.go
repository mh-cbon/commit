package main

import (
	"fmt"
	"os"

	"github.com/mh-cbon/go-repo-utils/repoutils"
	"github.com/mh-cbon/gen-version-file/GenVersionFile"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "commit"
	app.Version = GenVersionFile.Version()
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
	}

	app.Action = func(c *cli.Context) error {
		files := c.StringSlice("file")
		message := c.String("message")

		if len(files)==0 {
			cli.ShowAppHelp(c)
			return cli.NewExitError("Files are required", 1)
		}

		if len(message)==0 {
			cli.ShowAppHelp(c)
			return cli.NewExitError("Message is required", 1)
		}

    path, err := os.Getwd()
    exitWithError(err)

    vcs, err := repoutils.WhichVcs(path)
    exitWithError(err)

    sfile := make([]string, 0)
    for _, file := range files {
      sfile = append(sfile, string(file))
    	err = repoutils.Add(vcs, path, string(file))
      if err!=nil {
        fmt.Printf("Failed to add %s\n", file)
        // fmt.Println(err)
        // it does not exit on error, just print it
      }
    }

    err = repoutils.Commit(vcs, path, message, sfile)
    if err!=nil {
      fmt.Printf("Failed to commit %s\n", err)
      // fmt.Println(err)
      // it does not exit on error, just print it
    }

    os.Exit(0)
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
