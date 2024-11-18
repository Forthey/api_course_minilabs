package coins_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/dto"
	"server/tracer"
)

func GetTopCoinMarkets(id string) (dto.Markets, error) {
	response, err := http.Get(fmt.Sprintf("%s?id=%s", marketsApiPath, id))
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}

	var markets dto.Markets
	err = json.Unmarshal(body, &markets)
	if err != nil {
		return nil, fmt.Errorf(tracer.AppendTrace(tracer.GetTrace(), err.Error()))
	}

	return markets[:10], nil
}
