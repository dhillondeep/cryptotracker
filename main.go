package main

// date for the coins
type Date struct {
	Year int
	Month int
	Day int
	Time int
}

// coin information
type Coin struct {
	Name string
	Price int
	MarketCap int
	Volume int
	CirculatingSupply int
}

// total packet
type Information struct {
	Date Date
	Coin Coin
	News []string
}

func main() {
	return
}
