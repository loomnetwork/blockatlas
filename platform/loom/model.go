package loom

const loomAnnualRate = 20.00

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
