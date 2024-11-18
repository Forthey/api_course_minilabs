package dto

type MinimizedCoin struct {
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	NameID   string `json:"nameid"`
	Rank     uint8  `json:"rank"`
	PriceUSD string `json:"price_usd"`
}

type Coins []MinimizedCoin
