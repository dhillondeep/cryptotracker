package main

import cmc "github.com/miguelmota/go-coinmarketcap"

// struct for storing date
type Date struct {
    Year  int
    Month string
    Day   int
    Time  string
}

// struct for storing the whole coin information packet
type Information struct {
    Date Date
    Coin cmc.Coin
    News []string
}
