package monitor

import (
    "time"
    "log"

    . "cryptotracker/cmd/cryptotracker/helper"
    . "cryptotracker/cmd/cryptotracker/types"
)

// Execute the monitor command to run the program like a server and store crypto data to the repo(s)
func Execute(commits int, override bool) {
    configuration := CommonParsingAndValidation().(Configuration)

    interval := 1 * time.Hour
    duration := 1752000 * time.Hour // run the app for 200 years (means never stop lol)

    if configuration.Interval == "daily" {
        interval = 24 * time.Hour
    }

    // make the first commit now
    ExecuteGatherAndPush(commits, override, configuration)

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
