package beefy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAPYBoosts() (APYData, error) {
	apyBoostResp, err := http.Get(URL_APY_BOOSTS)
	if err != nil {
		return nil, fmt.Errorf("Error during request for APY boosts: %v", err)
	}
	defer apyBoostResp.Body.Close()

	apyBoostBody, err := io.ReadAll(apyBoostResp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error while reading response for APY boosts: %v", err)
	}

	var apyBoostData APYData
	err = json.Unmarshal(apyBoostBody, &apyBoostData)
	if err != nil {
		return nil, fmt.Errorf("Error while decoding JSON for APY boosts: %v", err)
	}

	return apyBoostData, nil
}
