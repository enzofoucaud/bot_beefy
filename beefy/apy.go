package beefy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APYValue struct {
	Value float64
	Valid bool
}

type APYData map[string]APYValue

func (a *APYValue) UnmarshalJSON(data []byte) error {
	var floatVal float64
	if err := json.Unmarshal(data, &floatVal); err == nil {
		a.Value = floatVal
		a.Valid = true
		return nil
	}

	var strVal string
	if err := json.Unmarshal(data, &strVal); err == nil {
		if strVal == "Infinity" {
			a.Value = 1e308
			a.Valid = true
		} else {
			a.Valid = false
		}
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into APYValue", data)
}

func GetAPY() (APYData, error) {
	apyResp, err := http.Get(URL_APY)
	if err != nil {
		return nil, fmt.Errorf("Error during request for APY: %v", err)
	}
	defer apyResp.Body.Close()

	apyBody, err := io.ReadAll(apyResp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error while reading response for APY: %v", err)
	}

	var apyData APYData
	err = json.Unmarshal(apyBody, &apyData)
	if err != nil {
		return nil, fmt.Errorf("Error while decoding JSON for APY: %v", err)
	}

	return apyData, nil
}
