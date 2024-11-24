package beefy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Structure pour contenir les informations des vaults
type Vault struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Token               string   `json:"token"`
	TokenAddress        string   `json:"tokenAddress"`
	TokenDecimals       int      `json:"tokenDecimals"`
	TokenProviderID     string   `json:"tokenProviderId"`
	EarnedToken         string   `json:"earnedToken"`
	EarnedTokenAddress  string   `json:"earnedTokenAddress"`
	EarnContractAddress string   `json:"earnContractAddress"`
	Oracle              string   `json:"oracle"`
	OracleID            string   `json:"oracleId"`
	Status              string   `json:"status"`
	PlatformID          string   `json:"platformId"`
	Assets              []string `json:"assets"`
	StrategyTypeID      string   `json:"strategyTypeId"`
	Risks               []string `json:"risks"`
	AddLiquidityUrl     string   `json:"addLiquidityUrl"`
	Network             string   `json:"network"`
	CreatedAt           int64    `json:"createdAt"`
	Chain               string   `json:"chain"`
	Strategy            string   `json:"strategy"`
	LastHarvest         int64    `json:"lastHarvest"`
	PricePerFullShare   string   `json:"pricePerFullShare"`
	APY                 float64  `json:"apy"`
	TVL                 float64  `json:"tvl"`

	// Move it next time
	IsCLM bool
}

func GetVaults() ([]Vault, error) {
	vaultsResp, err := http.Get(URL_VAULTS)
	if err != nil {
		return nil, fmt.Errorf("Error during request for vaults: %v", err)
	}
	defer vaultsResp.Body.Close()

	vaultsBody, err := io.ReadAll(vaultsResp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error while reading response for vaults: %v", err)
	}

	var vaults []Vault
	err = json.Unmarshal(vaultsBody, &vaults)
	if err != nil {
		return nil, fmt.Errorf("Error while decoding JSON for vaults: %v", err)
	}

	return vaults, nil
}
