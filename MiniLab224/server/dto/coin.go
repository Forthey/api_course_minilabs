package dto

import (
	"server/database/models"
)

type Coin struct {
	MinimizedCoin
	PercentChange24H string  `json:"percent_change_24h"`
	PercentChange1H  string  `json:"percent_change_1h"`
	PercentChange7D  string  `json:"percent_change_7d"`
	MarketCapUsd     string  `json:"market_cap_usd"`
	Volume24         float64 `json:"volume24"`
	Volume24Native   float64 `json:"volume24_native"`
	PriceBtc         string  `json:"price_btc"`
}

type EndorsedCoin struct {
	Coin
	models.EndorsedCoinInfo
}
