package loom

const loomAnnualRate = 20.00

var LockTimeTier = map[string]string{
	"TIER_ZERO":  "0", // two week
	"TIER_ONE":   "1", // three months
	"TIER_TWO":   "2", // six months
	"TIER_THREE": "3", // one year
}

var TierMap = map[uint64]string{
	0: "TIER_ZERO",
	1: "TIER_ONE",
	2: "TIER_TWO",
	3: "TIER_THREE",
}

var TierBonusMap = map[string]float64{
	"TIER_ZERO":  5.00,  // two weeks
	"TIER_ONE":   7.50,  // three months
	"TIER_TWO":   10.00, // six months
	"TIER_THREE": 20.00, // one year
}

// # Staking

type StakingPool struct {
	NotBondedTokens string `json:"not_bonded_tokens"`
	BondedTokens    string `json:"bonded_tokens"`
}

type ValidatorPage []Validator

type Validator struct {
	Address         string `json:"address,omitempty"`
	Jailed          bool   `json:"jailed,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	Image           string `json:"image,omitempty"`
	Website         string `json:"website,omitempty"`
	DelegationTotal string `json:"delegationTotal,omitempty"`
	Fee             string `json:"fee,omitempty"`
}
