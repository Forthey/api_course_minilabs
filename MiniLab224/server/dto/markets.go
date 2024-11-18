package dto

type Market struct {
	Name     string  `json:"name"`
	Base     string  `json:"base"`
	Quote    string  `json:"quote"`
	Price    float64 `json:"price"`
	PriceUsd float64 `json:"price_usd"`
}

type Markets []Market
