package helper

import (
    "time"
    "strconv"
    "path/filepath"
    "os"
    "runtime"
    "io/ioutil"

    cmc "github.com/miguelmota/go-coinmarketcap"
    "gopkg.in/yaml.v2"
    "os/exec"
    "math"
    "encoding/json"
    "strings"
    "log"
)

// Creates a coin information packet that contains the date, coin and news
func GetCoinInformation(coinName string) (Information) {
    return Information{
        Date: getDate(),
        Coin: getCoin(coinName),
    }
}

// Gets current date and converts it into struct which will be later on be used in JSON
func getDate() (Date) {
    currTime := time.Now()

    return Date{
        Year:  currTime.Year(),
        Month: currTime.Month().String(),
        Day:   currTime.Day(),
        Time: strconv.Itoa(currTime.Hour()) + ":" + strconv.Itoa(currTime.Minute()) + ":" +
            strconv.Itoa(currTime.Second()),
    }
}

// Gathers data from CoinMarketCap and returns the coin struct
// This struct will later be converted into JSON
func getCoin(name string) (cmc.Coin) {
    coinInfo, _ := cmc.GetCoinData(name)

    return coinInfo
}

// Convert File text to a String
func FileToString(fileName string) (string, error) {
    buff, err := ioutil.ReadFile(fileName)
    str := string(buff)

    return str, err
}

// Converts a String to Yml struct
func ToYmlStruct(data string, out interface{}) (error) {
    e := yaml.Unmarshal([]byte(data), out)

    return e
}

// Parses a YML file and convert the data to a struct
func Parse(fileName string) (interface{}) {
    c := Configuration{}
    data, _ := FileToString(fileName)
    ToYmlStruct(data, &c)

    return c
}

// Validates if the data in configuration file is valid
func Validate(config interface{}) (string, bool) {
    c := config.(Configuration)

    if len(c.Coins) == 0 {
        return "Error: Please add coins to track (assets/config.yml)", false
    }

    if len(c.Repos) == 0 {
        return "Error: Please add a repo to push data to (assets/config.yml)", false
    }

    if len(c.Interval) == 0 {
        c.Interval = "hourly"
        return "Warning: Hourly interval selected since no interval was specified (assets/config.yml)", true

    } else if c.Interval != "hourly" && c.Interval != "daily" {
        c.Interval = "hourly"
        return "Warning: Hourly interval selected since interval was invalid.", true
    }

    return "Successfully validated configuration file.", true
}

// Get's the path to the root directory of this project
func GetCryptotrackerPath() (string, error) {
    _, configFileName, _, _ := runtime.Caller(0)
    return filepath.Abs(configFileName + "/../../../../")
}

// Parses configuration file and validates data
func CommonParsingAndValidation() (interface{}) {
    log.Print("Parsing Configuration file")

    cryptoPath, _ := GetCryptotrackerPath()
    fullPath, _ := filepath.Abs(cryptoPath + "/assets/config/config.yml")

    log.Println("Validating Configuration file data")
    configuration := Parse(fullPath)
    message, valid := Validate(configuration)
    log.Println(message)

    if !valid {
        os.Exit(1)
    }

    return configuration
}

// Commits data to the git repo
func CommitData(path string, currCommits int) (int) {
    commitTime := strconv.Itoa(time.Now().Hour()) + ":" + strconv.Itoa(time.Now().Minute()) + ":" +
        strconv.Itoa(time.Now().Second())
    commitDate := time.Now().Month().String() + " " + strconv.Itoa(time.Now().Day()) + ", " +
        strconv.Itoa(time.Now().Year())

    currCommits++
    exec.Command("git", "-C", path, "commit", "-m", "pushed crypto data on " + commitDate + " @ time "+
        commitTime).Output()

    return currCommits
}

// Gathers data, add it to the repo, commits it and pushes it
func ExecuteGatherAndPush(commits int, override bool, configuration Configuration) {
    log.Println("Gathering Data and Pushing it to the repo(s): ")

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
                log.Println(r + " is not a valid coin. Skipping this!!")
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

    log.Println("Done!!")
}
