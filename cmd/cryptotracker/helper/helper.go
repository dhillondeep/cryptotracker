package helper

import (
    "time"
    "strconv"
    "fmt"
    "path/filepath"
    "os"
    "runtime"
    "io/ioutil"

    . "../types"
    cmc "github.com/miguelmota/go-coinmarketcap"
    "gopkg.in/yaml.v2"
)

// Creates a coin information packet that contains the date, coin and news
func getCoinInformation(coinName string) (Information) {
    return Information{
        Date: getDate(),
        Coin: getCoin(coinName),
    }
}

// Gets current date and converts it into struct which will be later on be
// used in JSON
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

func Parse(fileName string) (interface{}) {
    c := Configuration{}
    data, _ := FileToString(fileName)
    ToYmlStruct(data, &c)

    return c
}

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

func GetCryptotrackerPath() (string, error) {
    _, configFileName, _, _ := runtime.Caller(0)
    return filepath.Abs(configFileName + "/../../../../")
}

func CommonParsingAndValidation() (interface{}) {
    fmt.Print("Parsing Configuration file ... ")

    cryptoPath, _ := GetCryptotrackerPath()
    fullPath, _ := filepath.Abs(cryptoPath + "/assets/config/config.yml")

    fmt.Println("done")

    fmt.Println("Validating Configuration file data")
    configuration := Parse(fullPath)
    message, valid := Validate(configuration)
    fmt.Println(message)

    if !valid {
        os.Exit(1)
    }

    return configuration
}
