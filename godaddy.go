package main
import (
    "context"
    "log"
    "os"
    "github.com/urfave/cli"
    "github.com/Cabemo/godaddy-cli/internal/util"
    "github.com/Cabemo/godaddy-cli/internal/godaddygo"
)

var app = cli.NewApp()

func info() {
    app.Name = "Godaddy CLI"
    app.Usage = "A Godaddy CLI to manage domains"
    app.Author = "Cabemo"
    app.Version = "1.0.0"
}

func commands() {
    app.Commands = []cli.Command{
        {
            Name: "domains",
            Usage: "Manage your domains",
            Subcommands: []cli.Command{
                {
                    Name: "list",
                    Aliases: []string{"l","ls"},
                    Usage: "List your owned domains",
                    Action: func(c *cli.Context) {
                        godaddy, err := godaddygo.GetGoDaddy()

                        if err != nil {
                            panic(err.Error())
                        }

                        domains, err := godaddy.ListDomains(context.Background())

                        if err != nil {
                            panic(err.Error())
                        }

                        util.PrintDomains(domains)
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
