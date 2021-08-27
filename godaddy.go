package main

import (
	"context"
	"log"
	"os"

	"github.com/Cabemo/godaddy-cli/internal/config"
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
	flags := map[string]cli.Flag{
		"domain": &cli.StringFlag{
			Name:  "domain",
			Value: "",
			Usage: "The domain name to query e.g. example.com",
		},
	}
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
					Action: func(c *cli.Context) {
						godaddy, err := config.GoDaddy()

						if err != nil {
							util.PrintError(err)
						}

						domains, err := godaddy.ListDomains(context.Background())

						if err != nil {
							util.PrintError(err)
						}

						util.PrintDomains(domains)
					},
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
						flags["domain"],
					},
					Usage: "List the records of a domain",
					Action: func(c *cli.Context) {
						domainArg := c.String("domain")

						godaddy, err := config.GoDaddy()

						if err != nil {
							panic(err.Error())
						}

						domain := godaddy.Domain(domainArg)
						domain.GetDetails(context.Background())
						records := domain.Records()
						// Actual list of records
						recs, err := records.List(context.Background())

						if err != nil {
							util.PrintError(err)
						}

						util.PrintRecords(recs)
					},
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
