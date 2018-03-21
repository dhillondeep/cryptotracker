package gather

import (
    . "github.com/dhillondeep/cryptotracker/cmd/cryptotracker/helper"
    . "github.com/dhillondeep/cryptotracker/cmd/cryptotracker/types"
)

// Execute the gather command to run the program one time and store crypto data to the repo(s)
func Execute(commits int, override bool) {
    configuration := CommonParsingAndValidation().(Configuration)

    ExecuteGatherAndPush(commits, override, configuration)
}
