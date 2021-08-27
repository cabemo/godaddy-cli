// Package records has all the logic for managing records
package records

import (
	"context"

	"github.com/Cabemo/godaddy-cli/internal/config"
	"github.com/Cabemo/godaddy-cli/internal/util"
	"github.com/oze4/godaddygo"
	"github.com/urfave/cli"
)

// List function lists the records of a domain
func List(c *cli.Context) {
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
}

// Add adds a record to the specified domain
func Add(c *cli.Context) {
	domainArg := c.String("domain")
	recordType := util.StringToRecord(c.String("type"))
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

	util.PrintRecord(record, "Added")
}

// Remove deletes a specified record by oze4/godaddygo.RecordType and godaddygo.Record.Name
func Remove(c *cli.Context) {
	domainArg := c.String("domain")
	recordType := util.StringToRecord(c.String("type"))
	name := c.String("name")

	godaddy, err := config.GoDaddy()

	if err != nil {
		//TODO Handle errors on config and change on all methods used by cli
		panic(err.Error())
	}

	domain := godaddy.Domain(domainArg)
	domain.GetDetails(context.Background())
	records := domain.Records()

	record := godaddygo.Record{
		Type: recordType,
		Name: name,
	}
	err = records.Delete(context.Background(), record)

	if err != nil {
		util.PrintError(err)
		return
	}

	util.PrintRecord(record, "Deleted")
}
