package beefy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TVLData struct {
	ID []struct {
		Data map[string]float64
	}
}

func GetTVL() (TVLData, error) {
	resp, err := http.Get(URL_TVL)
	if err != nil {
		return TVLData{}, fmt.Errorf("Error during request for TVL: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TVLData{}, fmt.Errorf("Error while reading response for TVL: %v", err)
	}

	var tvlData TVLData
	err = json.Unmarshal(body, &tvlData)
	if err != nil {
		return TVLData{}, fmt.Errorf("Error while decoding JSON for TVL: %v", err)
	}

	return tvlData, nil
}
