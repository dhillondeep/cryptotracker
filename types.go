package main

import cmc "github.com/miguelmota/go-coinmarketcap"

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
