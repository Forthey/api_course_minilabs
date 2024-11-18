package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/coins_api"
)

const marketsPrefix = "/markets"

func getTopMarkets(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	markets, err := coins_api.GetTopCoinMarkets(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	response, err := json.Marshal(markets)
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
