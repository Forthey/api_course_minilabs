package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/coins_api"
)

const coinPrefix = "/coin"

func getCoinById(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	coin, err := coins_api.GetCoin(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	response, err := json.Marshal(coin)
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
