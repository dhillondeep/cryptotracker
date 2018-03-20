package gather

import (
    "os"
    "strings"
    "fmt"
    "os/exec"
    "math"

    . "../../helper"
    . "../../types"
    "path/filepath"
    "strconv"
    "time"
    "encoding/json"
    "io/ioutil"
)

func Execute(commits int, override bool) {
    // parse and validate configuration file
    configuration := CommonParsingAndValidation().(Configuration)
    fmt.Println("Gathering Data and Pushing it to the repo(s): ")

    allRepoPath, _ := GetCryptotrackerPath()
    allRepoPath += "/repositories"

    if _, err := os.Stat(allRepoPath); os.IsNotExist(err) {
        os.Mkdir(allRepoPath, os.ModePerm)
    } else {
        os.RemoveAll(allRepoPath)
        os.Mkdir(allRepoPath, os.ModePerm)
    }

    // make sure coins are unique
    coinMap := make(map[string]string)
    for i := 0; i < len(configuration.Coins); i++ {
        coinMap[configuration.Coins[i]] = configuration.Coins[i]
    }

    for i := 0; i < len(configuration.Repos); i++ {
        repoPath, _ := GetCryptotrackerPath()
        repoPath += "/repositories"
        splits := strings.Split(configuration.Repos[0], "/")
        repoPath += "/" + splits[len(splits)-1]
        os.Mkdir(repoPath, os.ModePerm)

        validCommits := int(math.Min(15, float64(len(configuration.Coins))))
        changes := false

        if commits < validCommits {
            validCommits = commits
        }
        currCommits := 1

        exec.Command("git", "clone", configuration.Repos[i], repoPath).Output()

        for r := range coinMap {
            info := GetCoinInformation(r)

            // check if coin is valid by looking at the data
            if info.Coin.Symbol == "" && info.Coin.Name == "" {
                fmt.Println(r + " is not a valid coin. Skipping this!!")
                validCommits -= 1
                continue
            }

            coinPath := repoPath + string(filepath.Separator) + r
            yearPath := coinPath + string(filepath.Separator) + strconv.Itoa(info.Date.Year)
            monthPath := yearPath + string(filepath.Separator) + info.Date.Month
            dayPath := monthPath + string(filepath.Separator) + strconv.Itoa(info.Date.Day)
            filePath := dayPath + string(filepath.Separator) + info.Coin.Symbol + "-hour" +
                strconv.Itoa(time.Now().Hour()) + ".json"

            if _, err := os.Stat(dayPath); os.IsNotExist(err) {
                os.MkdirAll(dayPath, os.ModePerm)
            }

            if _, err := os.Stat(filePath); os.IsNotExist(err) || override {
                coinJson, _ := json.MarshalIndent(info, "", " ")
                err = ioutil.WriteFile(filePath, coinJson, 0644)

                // add all the files
                exec.Command("git", "-C", repoPath, "add", "-A").Output()

                if currCommits < validCommits {
                    currCommits = CommitData(repoPath, currCommits)
                }

                changes = true
            }
        }

        if changes {
            // final commit and push
            CommitData(repoPath, currCommits)
            exec.Command("git", "-C", repoPath, "push", "origin", "master").Output()
        }
    }

    fmt.Println("Done!!")
}
