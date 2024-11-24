package beefy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APYBreakdownValue struct {
	ClmApr float64 `json:"clmApr"`
}

func GetAPYBreakdown() (map[string]APYBreakdownValue, error) {
	apyResp, err := http.Get(URL_APY_BREAKDOWN)
	if err != nil {
		return nil, fmt.Errorf("Error during request for APY: %v", err)
	}
	defer apyResp.Body.Close()

	apyBody, err := io.ReadAll(apyResp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error while reading response for APY: %v", err)
	}

	var data map[string]APYBreakdownValue
	err = json.Unmarshal(apyBody, &data)
	if err != nil {
		return nil, fmt.Errorf("Error while decoding JSON for APY: %v", err)
	}

	return data, nil
}
