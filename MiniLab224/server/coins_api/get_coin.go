package coins_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/dto"
	"server/tracer"
)

func GetCoin(id string) (*dto.Coin, error) {
	response, err := http.Get(fmt.Sprintf("%s?id=%s", coinApiPath, id))
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}

	var coin []dto.Coin
	err = json.Unmarshal(body, &coin)
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}

	if len(coin) == 0 {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), "Coin does not exist"))
	}

	return &coin[0], nil
}
