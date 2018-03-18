package main

import (
	cmc "github.com/miguelmota/go-coinmarketcap"
	"time"
	"strconv"
)

// date for the coins
type Date struct {
	Year  int
	Month string
	Day   int
	Time  string
}

// total packet
type Information struct {
	Date Date
	Coin cmc.Coin
	News []string
}

func main() {
	return
}

// Creates a coin information packet that contains the date, coin and news
func getCoinInformation(coinName string) (Information) {
	return Information{
		getDate(),
		getCoin(coinName),
		[]string{},
	}
}

// Gets current date and converts it into struct which will be later on be
// used in JSON
func getDate() (Date) {
	currTime := time.Now()

	return Date{
		currTime.Year(),
		currTime.Month().String(),
		currTime.Day(),
		strconv.Itoa(currTime.Hour()) + ":" + strconv.Itoa(currTime.Minute()) + ":" +
			strconv.Itoa(currTime.Second()),
	}
}

// Gathers data from CoinMarketCap and returns the coin struct
// This struct will later be converted into JSON
func getCoin(name string) (cmc.Coin) {
	coinInfo, _ := cmc.GetCoinData(name)

	return coinInfo
}
