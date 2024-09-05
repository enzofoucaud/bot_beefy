package safety

import "strings"

var riskScores = map[string]int{
	// Beefy Risks //
	// COMPLEXITY
	"COMPLEXITY_LOW":  1,
	"COMPLEXITY_MID":  2,
	"COMPLEXITY_HIGH": 3,
	// TIME IN MARKET
	"BATTLE_TESTED":      1,
	"NEW_STRAT":          2,
	"EXPERIMENTAL_STRAT": 3,
	// Asset Risks //
	// IMPERMANENT LOSS
	"IL_NONE":     1,
	"IL_LOW":      2,
	"IL_HIGH":     3,
	"ALGO_STABLE": 4,
	// LIQUIDITY
	"LIQ_HIGH": 1,
	"LIQ_LOW":  3,
	// MARKET CAP
	"MCAP_LARGE":  1,
	"MCAP_MEDIUM": 2,
	"MCAP_SMALL":  3,
	"MCAP_MICRO":  4,
	// SUPPLY
	"SUPPLY_CENTRALIZED": 3,
	// Platform Risks //
	// PLATFORM
	"PLATFORM_ESTABLISHED": 1,
	"PLATFORM_NEW":         3,
	// AUDIT
	"AUDIT":    1,
	"NO_AUDIT": 3,
	// CONTRACTS
	"CONTRACTS_VERIFIED":   1,
	"CONTRACTS_UNVERIFIED": 3,
	// ADMINS
	"ADMIN_WITH_TIMELOCK":    1,
	"ADMIN_WITHOUT_TIMELOCK": 3,
}

// Définir le seuil de sécurité
const safetyThreshold = 8

// CalculateSafetyScore calcule le score de sécurité pour une liste de risques donnée
func CalculateSafetyScore(risks []string) int {
	score := 0
	for _, risk := range risks {
		if val, exists := riskScores[strings.ToUpper(risk)]; exists {
			score += val
		}
	}
	return score
}

// IsPoolSafe vérifie si une pool est sûre en fonction de son score de sécurité
func IsPoolSafe(risks []string) bool {
	score := CalculateSafetyScore(risks)
	return score <= safetyThreshold
}
