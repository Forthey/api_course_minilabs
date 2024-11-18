package coins_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/dto"
	"server/tracer"
)

type info struct {
	CoinsNum uint64 `json:"coins_num"`
	Time     uint64 `json:"time"`
}

type getTopCoinsData struct {
	Coins dto.Coins `json:"data"`
	Info  info      `json:"info"`
}

func GetTopCoins(coinLimit, coinStart string) (dto.Coins, error) {
	response, err := http.Get(fmt.Sprintf("%s?start=%s&limit=%s", coinsApiPath, coinStart, coinLimit))
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}

	var coins getTopCoinsData
	err = json.Unmarshal(body, &coins)
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}

	return coins.Coins, nil
}
