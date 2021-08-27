package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Cabemo/godaddy-cli/internal/config"
	"github.com/Cabemo/godaddy-cli/internal/util"
	"github.com/oze4/godaddygo"
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
			Name:     "domain",
			Value:    "",
			Usage:    "The domain name to query e.g. example.com",
			Required: true,
		},
		"type": &cli.StringFlag{
			Name:     "type",
			Value:    "",
			Usage:    "The type of record, can be one of [A, AAAA, TXT, MX]",
			Required: true,
		},
		"name": &cli.StringFlag{
			Name:     "name",
			Value:    "",
			Usage:    "The name value of the record",
			Required: true,
		},
		"value": &cli.StringFlag{
			Name:     "value",
			Value:    "",
			Usage:    "The value of the record",
			Required: true,
		},
		"ttl": &cli.IntFlag{
			Name:  "ttl",
			Value: 600,
			Usage: "The TTL of the record (600 by default)",
		},
		"priority": &cli.IntFlag{
			Name:  "priority",
			Value: 0,
			Usage: "The priority when using an MX record",
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
				{
					Name:    "add",
					Aliases: []string{"a"},
					Flags: []cli.Flag{
						flags["domain"],
						flags["type"],
						flags["name"],
						flags["value"],
						flags["priority"],
						flags["ttl"],
					},
					Usage: "Add a record to the specified domain",
					Action: func(c *cli.Context) {
						domainArg := c.String("domain")
						rType := c.String("type")
						name := c.String("name")
						value := c.String("value")
						priority := c.Int("priority")
						ttl := c.Int("ttl")

						godaddy, err := config.GoDaddy()

						if err != nil {
							panic(err.Error())
						}

						domain := godaddy.Domain(domainArg)
						domain.GetDetails(context.Background())
						records := domain.Records()

						var recordType godaddygo.RecordType
						switch rType {
						case "A":
							recordType = godaddygo.RecordTypeA
							break
						case "AAAA":
							recordType = godaddygo.RecordTypeAAAA
							break
						case "CNAME":
							recordType = godaddygo.RecordTypeCNAME
							break
						case "MX":
							recordType = godaddygo.RecordTypeMX
							break
							//						case "NS":
							//							recordType = godaddygo.RecordTypeNS
							//							break
							//						case "SOA":
							//							recordType = godaddygo.RecordTypeSOA
							//							break
							//						case "SRV":
							//							recordType = godaddygo.RecordTypeSRV
							//							break
						case "TXT":
							recordType = godaddygo.RecordTypeTXT
							break
						default:
							panic(fmt.Sprintf("Invalid record type %s", rType))
						}

						record := godaddygo.Record{
							Type:     recordType,
							Name:     name,
							Data:     value,
							Priority: priority,
							TTL:      ttl,
						}
						err = records.Add(context.Background(), []godaddygo.Record{record})

						if err != nil {
							util.PrintError(err)
							return
						}

						fmt.Printf("Sucessfully added record: %s %s %s", record.Type, record.Name, record.Data)
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
