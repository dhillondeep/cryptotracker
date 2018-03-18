package main

import (
    "time"
    "os"
    "log"
    "fmt"

    "github.com/urfave/cli"
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
    app.UsageText = "Crypto Tracker - Track and gather cryptocurrency data"

    app.Commands = []cli.Command{
        {
            Name:      "gather",
            Usage:     "Runs the app one time, gathers data, and pushes it to repo(s)",
            UsageText: "gather - Runs the app one time, gathers data, and pushes it to repo(s)",
            Flags: []cli.Flag{
                cli.IntFlag{Name: "commits",
                    Usage: "Number of commits to make when data is pushed to repo(s) (5 is max)",
                    Value: 1},
                cli.IntFlag{Name: "news",
                    Usage: "Number of news articles to store (10 is max)",
                    Value: 1},
                cli.BoolFlag{Name: "override",
                    Usage: "Flag to override the data already in the repo (only if data is in supported form)"},
            },
            Before: func(c *cli.Context) error {
                fmt.Fprintf(c.App.Writer, "Cryptotracker working and gathering data for you!!\n")
                return nil
            },
            After: func(c *cli.Context) error {
                fmt.Fprintf(c.App.Writer, "Cryptotracker (gather) successfully finished\n")
                return nil
            },
            Action: func(c *cli.Context) error {
                // Code to run gather sub command goes here
                execute(false, c.Int("commits"), c.Int("news"), c.Bool("override"))
                return nil
            },
        },
        {
            Name:      "monitor",
            Usage:     "Runs the app like a server, continuously gathers data, and pushes it to repo(s)",
            UsageText: "monitor - Runs the app like a server, continuously gathers data, and pushes it to repo(s)",
            Flags: []cli.Flag{
                cli.IntFlag{Name: "commits",
                    Usage: "Number of commits to make when data is pushed to repo(s) (5 is max)",
                    Value: 1},
                cli.IntFlag{Name: "news",
                    Usage: "Number of news articles to store (10 is max)",
                    Value: 1},
            },
            Before: func(c *cli.Context) error {
                fmt.Fprintf(c.App.Writer, "Cryptotracker continuously working and gathering data for you!!\n")
                return nil
            },
            After: func(c *cli.Context) error {
                fmt.Fprintf(c.App.Writer, "Cryptotracker (monitor) successfully finished\n")
                return nil
            },
            Action: func(c *cli.Context) error {
                // Code to run monitor sub command goes here
                execute(true, c.Int("commits"), c.Int("news"), false)
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
