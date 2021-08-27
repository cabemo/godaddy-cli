package util

import "github.com/urfave/cli"

var Flags = map[string]cli.Flag{
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
