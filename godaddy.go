package main
import (
    "context"
    "fmt"
    "log"
    "os"
    "github.com/urfave/cli"
    "github.com/oze4/godaddygo"
    "github.com/Cabemo/godaddy-cli/internal/credentials"
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
                        creds, err := credentials.GetCredentials()
                        if err != nil {
                            panic(err.Error())
                        }

                        api, err := godaddygo.NewProduction(creds.Key, creds.Secret)
                        if err != nil {
                            panic(err.Error())
                        }
                        godaddy := api.V1()
                        domains, err := godaddy.ListDomains(context.Background())

                        if err != nil {
                            log.Fatal(err)
                        }

                        for _, domain := range domains {
                            fmt.Println(domain.Domain, domain.NameServers)
                        }
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
