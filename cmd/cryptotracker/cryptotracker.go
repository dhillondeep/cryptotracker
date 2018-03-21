package main

import (
    "time"
    "os"
    "log"

    "github.com/urfave/cli"
    "cryptotracker/cmd/cryptotracker/commands/gather"
    "cryptotracker/cmd/cryptotracker/commands/monitor"
)

func main() {
    app := cli.NewApp()
    app.Name = "Cryptotracker"
    app.Version = "0.0.1"
    app.EnableBashCompletion = true
    app.Compiled = time.Now()
    app.Authors = []cli.Author{
        {
            Name:  "Deep Dhillon",
            Email: "deep@deepdhillon.ca",
        },
    }
    app.Copyright = "Copyright (c) [2018] [Deep Dhillon]"
    app.Usage = "Track and gather cryptocurrency data"
    app.UsageText = "cryptotracker [global options] command"

    app.Commands = []cli.Command{
        {
            Name:      "gather",
            Usage:     "Runs the app one time, gathers data, and pushes it to repo(s)",
            UsageText: "gather - Runs the app one time, gathers data, and pushes it to repo(s)",
            Flags: []cli.Flag{
                cli.IntFlag{Name: "commits",
                    Usage: "Number of commits to make when data is pushed to repo(s) (max = min(num coins, 15))",
                    Value: 1},
                cli.BoolFlag{Name: "override",
                    Usage: "Flag to override the data already in the repo (only if data is in supported form)"},
            },
            Action: func(c *cli.Context) error {
                // Code to run gather sub command goes here
                gather.Execute(c.Int("commits"), c.Bool("override"))
                return nil
            },
        },
        {
            Name:      "monitor",
            Usage:     "Runs the app like a server, continuously gathers data, and pushes it to repo(s)",
            UsageText: "monitor - Runs the app like a server, continuously gathers data, and pushes it to repo(s)",
            Flags: []cli.Flag{
                cli.IntFlag{Name: "commits",
                    Usage: "Number of commits to make when data is pushed to repo(s) (max = min(num coins, 15))",
                    Value: 1},
                cli.BoolFlag{Name: "override",
                    Usage: "Flag to override the data already in the repo (only if data is in supported form)"},
            },
            Action: func(c *cli.Context) error {
                // Code to run monitor sub command goes here
                monitor.Execute(c.Int("commits"), c.Bool("override"))
                return nil
            },
        },
    }

    app.Action = func(c *cli.Context) error {
        app.Command("help").Run(c)
        return nil
    }

    err := app.Run(os.Args)

    if err != nil {
        log.Fatal(err)
    }
}
