package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/coins_api"
)

const coinsPrefix = "/coins"

func getCoins(writer http.ResponseWriter, request *http.Request) {
	limit, start := request.URL.Query().Get("limit"), request.URL.Query().Get("start")
	if limit == "" || start == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	coins, err := coins_api.GetTopCoins(limit, start)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	response, err := json.Marshal(coins)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	_, err = fmt.Fprintln(writer, string(response))
	if err != nil {
		fmt.Println(err.Error())
	}
}
