package gather

import (
    "os"
    "strings"
    "fmt"
    "os/exec"
    "math"

    . "../../helper"
    . "../../types"
)

func Execute(commits int, override bool) {
    // parse and validate configuration file
    configuration := CommonParsingAndValidation().(Configuration)
    commits = int(math.Min(float64(len(configuration.Coins)), 15))

    fmt.Println("Gathering Data and Pushing it to the repo(s)")

    allRepoPath, _ := GetCryptotrackerPath()
    allRepoPath += "/repositories"

    if _, err := os.Stat(allRepoPath); os.IsNotExist(err) {
        os.Mkdir(allRepoPath, os.ModePerm)
    } else {
        os.RemoveAll(allRepoPath)
        os.Mkdir(allRepoPath, os.ModePerm)
    }

    for i := 0; i < len(configuration.Repos); i++ {
        repoPath, _ := GetCryptotrackerPath()
        repoPath += "/repositories"
        splits := strings.Split(configuration.Repos[0], "/")
        repoPath += "/" + splits[len(splits)-1]
        os.Mkdir(repoPath, os.ModePerm)

        exec.Command("git", "clone", configuration.Repos[i], repoPath).Output()
    }
}
