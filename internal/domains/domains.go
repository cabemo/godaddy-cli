// Package domains has all logic to manage domains
package domains

import (
	"context"

	"github.com/Cabemo/godaddy-cli/internal/config"
	"github.com/Cabemo/godaddy-cli/internal/util"
	"github.com/urfave/cli"
)

// List lists all of the domains that the user has
func List(c *cli.Context) {
	godaddy, err := config.GoDaddy()

	if err != nil {
		util.PrintError(err)
	}

	domains, err := godaddy.ListDomains(context.Background())

	if err != nil {
		util.PrintError(err)
	}

	util.PrintDomains(domains)
}
