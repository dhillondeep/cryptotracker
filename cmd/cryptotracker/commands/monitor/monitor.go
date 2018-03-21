package monitor

import (
    "time"
    "log"

    . "github.com/dhillondeep/cryptotracker/cmd/cryptotracker/helper"
    . "github.com/dhillondeep/cryptotracker/cmd/cryptotracker/types"
)

// Execute the monitor command to run the program like a server and store crypto data to the repo(s)
func Execute(commits int, override bool) {
    configuration := CommonParsingAndValidation().(Configuration)

    interval := 1 * time.Minute
    duration := 1752000 * time.Hour // run the app for 200 years (means never stop lol)

    if configuration.Interval == "daily" {
        interval = 24 * time.Hour
    }

    ticker := time.NewTicker(interval)
    quit := make(chan struct{})
    go func() {
        for {
            select {
            case <-ticker.C:
                ExecuteGatherAndPush(commits, override, configuration)
            case <-quit:
                ticker.Stop()
                return
            }
        }
    }()

    log.Printf("Gathering and storing data each %v during %v.", interval, duration)
    <-time.After(time.Duration(duration))
    log.Println("Finished Execution")
}
