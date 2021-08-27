package main

import (
	"log"
	"os"

	"github.com/Cabemo/godaddy-cli/internal/domains"
	"github.com/Cabemo/godaddy-cli/internal/records"
	"github.com/Cabemo/godaddy-cli/internal/util"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "GoDaddy CLI"
	app.Usage = "A GoDaddy CLI to manage domains"
	app.Author = "Cabemo"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "domains",
			Aliases: []string{"d"},
			Usage:   "Manage your GoDaddy domains",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l", "ls"},
					Usage:   "List your owned domains",
					Action:  domains.List,
				},
			},
		},
		{
			Name:    "records",
			Aliases: []string{"r"},
			Usage:   "Manage the records of a domain",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls", "l"},
					Flags: []cli.Flag{
						util.Flags["domain"],
					},
					Usage:  "List the records of a domain",
					Action: records.List,
				},
				{
					Name:    "add",
					Aliases: []string{"a"},
					Flags: []cli.Flag{
						util.Flags["domain"],
						util.Flags["type"],
						util.Flags["name"],
						util.Flags["value"],
						util.Flags["priority"],
						util.Flags["ttl"],
					},
					Usage:  "Add a record to the specified domain",
					Action: records.Add,
				},
				{
					Name:    "remove",
					Aliases: []string{"rm", "r"},
					Flags: []cli.Flag{
						util.Flags["domain"],
						util.Flags["type"],
						util.Flags["name"],
					},
					Usage:  "Delete a record of the specified domain",
					Action: records.Remove,
				},
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
