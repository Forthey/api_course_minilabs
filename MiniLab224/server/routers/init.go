package routers

import "net/http"

func Init(mux *http.ServeMux) {
	mux.HandleFunc(coinPrefix, getCoinById)
	mux.HandleFunc(coinsPrefix, getCoins)
	mux.HandleFunc(marketsPrefix, getTopMarkets)
	mux.HandleFunc(endorsePrefix, endorsement)
}
