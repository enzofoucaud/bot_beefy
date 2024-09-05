package main

import (
	"bot_beefy/beefy"
	"fmt"
	"sort"
	"strings"
)

var tokensToVoid = map[string]bool{
	"ETHFI": true,
}

// Fonction pour filtrer et trier les vaults selon les mots clés
func filterVaults(vaults []beefy.Vault, keyword string) []beefy.Vault {
	var filteredVaults []beefy.Vault

	for _, vault := range vaults {
		if vault.Status == "eol" {
			continue
		}

		count := 0
		for _, asset := range vault.Assets {
			if tokensToVoid[asset] {
				count = 0
				break
			}
			if strings.Contains(asset, keyword) {
				count++
			}
		}
		if count >= 2 && count == len(vault.Assets) {
			filteredVaults = append(filteredVaults, vault)
		}
	}

	sort.Slice(filteredVaults, func(i, j int) bool {
		return filteredVaults[i].APY > filteredVaults[j].APY
	})

	return filteredVaults
}

func getVaults() []beefy.Vault {
	vaults, err := beefy.GetVaults()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	apyData, err := beefy.GetAPY()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	apyBoostData, err := beefy.GetAPYBoosts()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i, vault := range vaults {
		if apy, exists := apyData[vault.ID]; exists && apy.Valid {
			vaults[i].APY = apy.Value
			if boost, exists := apyBoostData["moo_"+vault.ID]; exists && boost.Valid {
				vaults[i].APY += boost.Value
			}
		} else {
			vaults[i].APY = 0.0
		}
	}

	return vaults
}
